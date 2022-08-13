package plan

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/internal/generated"
	"github.com/aakarim/pland/cli/internal/graphclient"
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
