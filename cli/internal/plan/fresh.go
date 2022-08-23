package plan

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/internal/generated"
	"github.com/aakarim/pland/cli/internal/graphclient"
	planEntity "github.com/aakarim/pland/pkg/plan"
	"github.com/charmbracelet/charm/client"
)

type PlanService struct {
	cfg *config.Config
}

func NewPlanService(cfg *config.Config) *PlanService {
	return &PlanService{cfg: cfg}
}

const (
	_                       = iota
	StatusFinishedWithClean = iota
	StatusFinishedWithCopy  = iota
)

func (p *PlanService) Fresh() error {
	// load the existing plan file
	homePlan, err := p.GetLocalPlan()
	if err != nil {
		return err
	}
	defer homePlan.Close()
	planB, err := io.ReadAll(homePlan)
	if err != nil {
		return fmt.Errorf("io.ReadAll(): %w", err)
	}
	homePlan.Close()
	planFileStr := string(planB)
	// validate
	if err := Validate(planFileStr); err != nil {
		return err
	}
	// parse
	pl, err := planEntity.Parse(context.Background(), planFileStr)
	if err != nil {
		return fmt.Errorf("planEntity.Parse(): %w", err)
	}

	// add fun header
	if pl.Header.Contents == "" {
		pl.Header.Contents = "			🕴🏼			"
	}

	beginningOfDay := time.Now().Truncate(24 * time.Hour)

	newEntry := planEntity.Day{
		Contents: `- [ ] share what you're working on`,
		Date:     beginningOfDay,
	}
	// get most recent date
	if len(pl.Days) > 0 {
		latestDate := pl.Days[0].Date
		if latestDate.After(beginningOfDay) || latestDate.Equal(beginningOfDay) {
			return fmt.Errorf("latest date in plan is %v, this is at the same time or after the current day so there's no need to run `fresh`. If you've flown to another timezone please take a nap.", latestDate)
		}
	}

	// add a new entry at the beginning of the list
	pl.Days = append([]planEntity.Day{newEntry}, pl.Days...)
	if err := p.OverwriteLocalPlan(pl); err != nil {
		return err
	}
	return nil
}

func (p *PlanService) OverwriteLocalPlan(pl *planEntity.PlanFile) error {
	// overwrite file with new version
	homePlanPath, err := p.GetLocalPlanPath()
	if err != nil {
		return err
	}

	newPlan, err := os.Create(homePlanPath)
	if err != nil {
		return fmt.Errorf("os.Create(%s): %w", homePlanPath, err)
	}
	defer newPlan.Close()
	plStr := pl.String()
	if _, err := newPlan.Write([]byte(plStr)); err != nil {
		return fmt.Errorf("file.Write(%s); %w", plStr[100:]+"...", err) // print out 100 first characters
	}
	return nil
}

func (p *PlanService) GetLocalPlanPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	homePlanPath := filepath.Join(home, ".plan")
	return homePlanPath, nil
}

func (p *PlanService) GetLocalPlan() (*os.File, error) {
	homePlanPath, err := p.GetLocalPlanPath()
	if err != nil {
		return nil, err
	}

	homePlan, err := os.Open(homePlanPath)
	if err != nil {
		return nil, fmt.Errorf("Open(%s): %w", homePlanPath, err)
	}
	return homePlan, err
}

func (p *PlanService) CopyOrCreatePlanFromStore(targetPath string) (int, error) {
	// first check if we have an active plan file in the store
	c, err := client.NewClientWithDefaults()
	if err != nil {
		return 0, err
	}
	j, err := c.JWT("charm")
	if err != nil {
		return 0, err
	}
	u, err := url.Parse(fmt.Sprintf("%s://%s:%d%s", p.cfg.Server.HttpScheme, p.cfg.Server.Host, p.cfg.Server.GraphQLPort, p.cfg.Server.GraphQLPath))
	if err != nil {
		return 0, err
	}
	httpClient := &http.Client{
		Transport: graphclient.NewAuthedTransport(j, http.DefaultTransport),
	}
	gqlClient := graphql.NewClient(u.String(), httpClient)
	resp, err := generated.GetLatestPlan(context.Background(), gqlClient)
	if err != nil {
		return 0, err
	}
	// copy over or create
	newFile, err := os.Create(targetPath)
	if err != nil {
		return 0, err
	}
	defer newFile.Close()

	latestPlan := resp.Me.GetPlan()
	if latestPlan.Id == "" {
		return StatusFinishedWithClean, nil
	}

	_, err = newFile.Write([]byte(latestPlan.Txt))
	if err != nil {
		return 0, err
	}
	return StatusFinishedWithCopy, nil
}
