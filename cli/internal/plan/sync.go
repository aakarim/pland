package plan

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/aakarim/pland/cli/internal/actions"
	"github.com/aakarim/pland/cli/internal/generated"
	"github.com/aakarim/pland/cli/internal/graphclient"
	"github.com/aakarim/pland/cli/internal/store"
	planEntity "github.com/aakarim/pland/pkg/plan"

	"github.com/charmbracelet/charm/client"
)

func (p *PlanService) Sync() (conflict bool, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return conflict, err
	}

	// find current plan and copy it over into store
	// if current plan is old then wipe it (mentioning to user that we've done this)
	homePlanPath := filepath.Join(home, ".plan")
	homePlan, err := planFromPath(homePlanPath)
	if err != nil {
		return conflict, fmt.Errorf("planFromPath(%s): %w", homePlanPath, err)
	}

	// publish the current plan
	c, err := client.NewClientWithDefaults()
	if err != nil {
		return conflict, actions.AuthError{Err: err}
	}
	j, err := c.JWT("charm")
	if err != nil {
		return conflict, err
	}
	u, err := url.Parse(fmt.Sprintf("%s://%s:%d%s", p.cfg.Server.HttpScheme, p.cfg.Server.Host, p.cfg.Server.GraphQLPort, p.cfg.Server.GraphQLPath))
	if err != nil {
		return conflict, err
	}
	httpClient := &http.Client{
		Transport: graphclient.NewAuthedTransport(j, http.DefaultTransport),
	}
	gqlClient := graphql.NewClient(u.String(), httpClient)

	resp, err := generated.CreatePlan(context.Background(), gqlClient, homePlan.String(), time.Now())
	if err != nil {
		return conflict, fmt.Errorf("error creating: %w", err)
	}
	conflict = resp.CreatePlan.HasConflict
	// if the created plan's parent is different then we hae a new version and overwrite
	// the digest includes the parent, so it will be diferent if parents are different
	if resp.CreatePlan.Digest != homePlan.Digest() {
		log.Println("local file is old, replacing with server version...")
		// write to an 'old' file first as backup
		overwriteF, err := os.Create(homePlanPath)
		if err != nil {
			return conflict, err
		}
		if _, err := overwriteF.Write([]byte(resp.CreatePlan.Txt)); err != nil {
			overwriteF.Close()
			return conflict, fmt.Errorf("file.Write(%s): %w", homePlanPath, err)
		}
		overwriteF.Close()
		log.Println("overwritten")
	}
	log.Println("synced", resp.CreatePlan.Id)
	return conflict, err
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
func planFromPath(srcPath string) (*planEntity.PlanFile, error) {
	f, err := os.Open(srcPath)
	if err != nil {
		return nil, fmt.Errorf("Open(%s): %w", srcPath, err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("ReadAll(): %w", err)
	}
	return planEntity.Parse(context.Background(), string(b))
}
