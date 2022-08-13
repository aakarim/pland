package sync

import (
	"fmt"

	"github.com/aakarim/pland/cli/internal/actions"
	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/internal/plan"
	ourCommon "github.com/aakarim/pland/cli/ui/common"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/charm/ui/common"
)

type errMsg error

var (
	styles    = common.DefaultStyles()
	paragraph = styles.Paragraph.Render
	keyword   = styles.Keyword.Render
	code      = styles.Code.Render
	subtle    = styles.Subtle.Render
)

const (
	statusInit              = iota
	statusFinishedWithClean = iota
	statusFinishedWithCopy  = iota
)

type model struct {
	ourCommon.Model
	hasUser     bool
	userName    string
	syncStatus  int
	spinner     spinner.Model
	status      int
	planPath    string
	planService *plan.PlanService
	cfg         *config.Config
}

func InitialModel(planService *plan.PlanService, cfg *config.Config) model {
	return model{hasUser: true, spinner: common.NewSpinner(), planService: planService, cfg: cfg}
}

const (
	_              = iota
	SyncPending    = iota
	SyncInProgress = iota
	SyncCompleted  = iota
)

func (m model) Init() tea.Cmd {
	return tea.Batch(
		actions.CheckUser,
		spinner.Tick,
	)
}

type userUpdated struct{}
type syncCompleted struct {
}

type freshPlan struct{ path string }
type freshPlanCreated struct {
	path   string
	status int
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	authM, authMsg := actions.UpdateAuth(m, msg)
	if authM != nil {
		return authM, authMsg
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			m.Quitting = true
			return m, tea.Quit
		}

	case errMsg:
		m.Err = msg
		fmt.Println("rendering error")
		return m, tea.Quit

	case userUpdated:
		m.hasUser = true
		return m, m.sync
	case freshPlan:
		m.planPath = msg.path
		return m, m.fresh
	case freshPlanCreated:
		m.status = msg.status
		return m, tea.Quit
	case actions.UserFound:
		m.hasUser = true
		m.userName = msg.Name
		m.syncStatus = SyncInProgress
		return m, m.sync
	case syncCompleted:
		m.syncStatus = SyncCompleted
		return m, tea.Quit
	}
	return m, cmd

}

func (m model) View() string {
	if m.Err != nil {
		fmt.Println(m.Err)
		return styles.Error.Render(m.Err.Error()) + "\n"
	}
	var str string
	if m.userName != "" {
		str += paragraph(fmt.Sprintf("logged in as %s\n\n", m.userName))
	}
	if m.status == statusFinishedWithClean {
		str += paragraph(fmt.Sprintf("plan file created at %s. Add your plan for the day there and re-run `plan` to publish.", ourCommon.Code(m.planPath)))
	}
	if m.status == statusFinishedWithCopy {
		str += paragraph(fmt.Sprintf("plan file copied from store to %s.", ourCommon.Code(m.planPath)))
	}
	if m.syncStatus == SyncInProgress {
		str += paragraph(fmt.Sprintf("\n\nsyncing...\n\n"))
	}
	if m.syncStatus == SyncCompleted {
		str += paragraph(fmt.Sprintf("sync completed.\n"))
	}
	if m.Quitting {
		return str + "\n"
	}
	return str
}
