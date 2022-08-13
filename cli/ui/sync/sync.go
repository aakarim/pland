package sync

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/aakarim/pland/cli/internal/actions"
	"github.com/aakarim/pland/cli/internal/generated"
	"github.com/aakarim/pland/cli/internal/graphclient"
	"github.com/aakarim/pland/cli/internal/store"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/charm/client"
)

type PlanDay string

type DayPlan struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	TS   int64     `json:"timestamp"`
	Hash uint64    `json:"hash"`
	Txt  string    `json:"txt"`
	Date time.Time `json:"date"`
}

func printFormatted(s ...string) {
	for _, ss := range s {
		fmt.Printf(paragraph(ss))
	}
	fmt.Println()
}

func (m model) sync() tea.Msg {
	rightNow := time.Now()
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	planPath := path.Join(home, ".goplan")

	// create goplan path
	if _, err := os.Open(planPath); err != nil && errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(planPath, os.ModeAppend); err != nil {
			return fmt.Errorf("os.Mkdir(%s): %w", planPath, err)
		}
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
	homePlan.Close()

	// get latest plan
	c, err := client.NewClientWithDefaults()
	if err != nil {
		return actions.AuthError{Err: err}
	}
	j, err := c.JWT("charm")
	if err != nil {
		return err
	}
	u, err := url.Parse(fmt.Sprintf("%s://%s:%d%s", m.cfg.Server.HttpScheme, m.cfg.Server.Host, m.cfg.Server.GraphQLPort, m.cfg.Server.GraphQLPath))
	if err != nil {
		return err
	}
	httpClient := &http.Client{
		Transport: graphclient.NewAuthedTransport(j, http.DefaultTransport),
	}
	gqlClient := graphql.NewClient(u.String(), httpClient)
	latestPlanResp, err := generated.GetLatestPlan(context.Background(), gqlClient)
	if err != nil {
		return err
	}

	// copy to database producing
	var currentPlanPath string
	currentPlanPath, err = copyPlanToDir(planPath, homePlanPath)
	if err != nil {
		return fmt.Errorf("copyPlanToDir(%s, %s): %w", planPath, homePlanPath, err)
	}
	// open new plan
	todaysPlan, err = planFromPath(currentPlanPath)
	if err != nil {
		return fmt.Errorf("planFromPath(%s): %w", currentPlanPath, err)
	}
	// publish the current plan
	if todaysPlan != nil {
		httpClient := &http.Client{
			Transport: graphclient.NewAuthedTransport(j, http.DefaultTransport),
		}
		gqlClient := graphql.NewClient(u.String(), httpClient)
		resp, err := generated.CreatePlan(context.Background(), gqlClient, todaysPlan.Txt, todaysPlan.Date)
		if err != nil {
			return fmt.Errorf("error creating: %w", err)
		}
		// if the created plan is older than the latest plan then overwrite local plan file
		if latestPlanResp.Me.Plan.Timestamp.After(resp.CreatePlan.Timestamp) {
			fmt.Println("local file is old, replacing with server version...")
			overwriteF, err := os.Create(homePlanPath)
			if err != nil {
				return err
			}
			if _, err := overwriteF.Write([]byte(latestPlanResp.Me.Plan.Txt)); err != nil {
				overwriteF.Close()
				return fmt.Errorf("file.Write(%s): %w", homePlanPath, err)
			}
			overwriteF.Close()
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
	return syncCompleted{}
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
