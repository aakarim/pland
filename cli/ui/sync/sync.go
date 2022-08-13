package sync

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/aakarim/pland/cli/internal/actions"
	"github.com/aakarim/pland/cli/internal/generated"
	"github.com/aakarim/pland/cli/internal/graphclient"
	"github.com/aakarim/pland/cli/internal/store"
	"github.com/cespare/xxhash"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/charm/client"
	"github.com/charmbracelet/charm/kv"
	"github.com/dgraph-io/badger/v3"
)

type PlanDay string

type planStore map[PlanDay]DayPlan

type DayPlan struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	TS   int64     `json:"timestamp"`
	Hash uint64    `json:"hash"`
	Txt  string    `json:"txt"`
	Date time.Time `json:"date"`
}

func syncExec() tea.Cmd {
	return func() tea.Msg {
		return sync()
	}
}
func printFormatted(s ...string) {
	for _, ss := range s {
		fmt.Printf(paragraph(ss))
	}
	fmt.Println()
}

func sync() tea.Msg {
	rightNow := time.Now()
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	planPath := path.Join(home, ".goplan")

	// Open a database (or create one if it doesn’t exist)
	db, err := kv.OpenWithDefaults("goplan")
	if err != nil {
		return err
	}
	defer db.Close()

	// Fetch updates and easily define your own syncing strategy
	if err := db.Sync(); err != nil {
		return err
	}

	// load all entries
	var dirtyStore bool
	allB, err := db.Get([]byte("plans"))
	if err != nil && !errors.Is(err, badger.ErrKeyNotFound) {
		return err
	}
	ps := planStore{}
	if !errors.Is(err, badger.ErrKeyNotFound) {
		if err := json.Unmarshal(allB, &ps); err != nil {
			return err
		}
	} else {
		dirtyStore = true
	}

	var todaysPlan *Plan
	// find current plan and copy it over into store
	// if current plan is old then wipe it (mentioning to user that we've done this)
	homePlanPath := filepath.Join(home, ".plan")
	homePlan, err := os.Open(homePlanPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("Open(%s): %w", homePlanPath, err)
	}
	// create file and exit
	if errors.Is(err, os.ErrNotExist) {
		return freshPlan{path: homePlanPath}
	}

	// copy to database producing
	var currentPlanPath string
	if homePlan != nil {
		currentPlanPath, err = copyPlanToDir(planPath, homePlanPath)
		if err != nil {
			return fmt.Errorf("copyPlanToDir(%s, %s): %w", planPath, homePlanPath, err)
		}
		// open new plan
		todaysPlan, err = planFromPath(currentPlanPath)
		if err != nil {
			return fmt.Errorf("planFromPath(%s): %w", currentPlanPath, err)
		}
	}

	// publish the current plan
	if todaysPlan != nil {
		c, err := client.NewClientWithDefaults()
		if err != nil {
			return actions.AuthError{Err: err}
		}
		j, err := c.JWT("charm")
		if err != nil {
			return err
		}
		u, err := url.Parse("http://localhost:8080/query")
		if err != nil {
			return err
		}
		httpClient := &http.Client{
			Transport: graphclient.NewAuthedTransport(j, http.DefaultTransport),
		}
		gqlClient := graphql.NewClient(u.String(), httpClient)
		resp, err := generated.CreatePlan(context.Background(), gqlClient, todaysPlan.Txt, todaysPlan.Date)
		if err != nil {
			return fmt.Errorf("error creating: %w", err)
		}
		fmt.Println("synced", resp.CreatePlan.Id)
	}

	// if the active plan is older than today then wipe the plan
	currentPlanDate, err := store.GetDateFromFilePath(currentPlanPath)
	if err != nil {
		return fmt.Errorf("store.GetDateFromFilePath(%s): %w", currentPlanDate, err)
	}

	beginningOfDay := rightNow.Truncate(24 * time.Hour)
	if beginningOfDay.After(currentPlanDate) {
		fmt.Println("wiping existing plan file")
		f, err := os.Create(homePlanPath)
		if err != nil {
			return fmt.Errorf("os.Create(%s): %w", homePlanPath, err)
		}
		f.Close()
	}

	// update managed store
	if err := filepath.WalkDir(planPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		extension := filepath.Ext(path)
		dir := filepath.Dir(path)
		if extension == ".txt" || extension == ".plan" {
			// get date from filename
			base := filepath.Base(path)
			// rename todo to include date
			p := path
			if base == "todo.txt" || base == ".plan" {
				newPath, err := copyPlanToDir(dir, p)
				if err != nil {
					return fmt.Errorf("copyPlanToDir(%s, %s): %w", dir, p, err)
				}
				p = newPath
				dirtyStore = true
			}
			fileDate := PlanDay(strings.TrimPrefix(strings.TrimSuffix(filepath.Base(p), filepath.Ext(p)), "plan-"))
			f, err := os.Open(p)
			defer f.Close()
			if err != nil {
				return err
			}
			b, err := io.ReadAll(f)
			if err != nil {
				return err
			}
			currentTxt := string(b)
			// get existing one from store
			var foundDate DayPlan
			for date, v := range ps {
				if date == fileDate {
					foundDate = v
				}
			}
			// get the most recent timestamp (that allows us to undo and save)
			var mostRecent Plan
			for _, pd := range foundDate.Plans {
				if pd.TS > mostRecent.TS {
					mostRecent = pd
				}
			}
			sum := xxhash.Sum64String(currentTxt)
			if sum != mostRecent.Hash {
				currentP := Plan{
					TS:   rightNow.Unix(),
					Hash: sum,
					Txt:  currentTxt,
				}
				foundDate.Plans = append(foundDate.Plans, currentP)
				ps[fileDate] = foundDate
				dirtyStore = true
			}
		}
		return nil
	}); err != nil {
		return err
	}

	if dirtyStore {
		jsonB, err := json.Marshal(ps)
		if err != nil {
			return err
		}
		if err := db.Set([]byte("plans"), jsonB); err != nil {
			return err
		}
	}
	return syncCompleted{
		store: ps,
	}
}

// copyPlanToDir copies the plan to the managed directory and gives it a suffix.
// The suffix date string is based on the last touched file date.
//
// returns the path to the destination file.
func copyPlanToDir(dstDir string, srcPath string) (string, error) {
	src, err := os.Open(srcPath)
	if err != nil {
		return "", fmt.Errorf("Open(%s): %w", srcPath, err)
	}
	fileInfo, err := src.Stat()
	if err != nil {
		return "", fmt.Errorf("Stat(): %w", err)
	}
	t := fileInfo.ModTime()
	dstPath := filepath.Join(dstDir, fmt.Sprintf("plan-%s.txt", store.GetFileDateSuffix(t)))
	newF, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("Open(%s): %w", dstPath, err)
	}
	defer newF.Close()

	_, err = io.Copy(newF, src)
	if err != nil {
		return "", fmt.Errorf("Copy(): %w", err)
	}
	return dstPath, nil
}

// planFromPath creates a plan from the path's contents
// does not set the timestamp or hash because the timestamp is ambiguous until
// it is sent to the server.
// uses date from the file's path
func planFromPath(srcPath string) (*Plan, error) {
	d, err := store.GetDateFromFilePath(srcPath)
	if err != nil {
		return nil, fmt.Errorf("store.GetDateFromFilePath(%s): %w", srcPath, err)
	}
	f, err := os.Open(srcPath)
	if err != nil {
		return nil, fmt.Errorf("Open(%s): %w", srcPath, err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("ReadAll(): %w", err)
	}
	return &Plan{
		Txt:  string(b),
		Date: d,
	}, nil
}
