package fresh

import (
	"errors"
	"os"

	"github.com/aakarim/pland/cli/internal/actions"
	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/internal/plan"
	"github.com/aakarim/pland/cli/ui/common"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	common.Model
	planService *plan.PlanService
	cfg         *config.Config
	status      int
}
type errMsg error
type success struct{}

const (
	statusSuccess = 1 + iota
)

func InitialModel(planService *plan.PlanService, cfg *config.Config) model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return m.fresh
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	authM, authMsg := actions.UpdateAuth(m, msg)
	if authM != nil {
		return authM, authMsg
	}
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case success:
		m.status = statusSuccess
		m.Quitting = true
		return m, tea.Quit
	case errMsg:
		m.Err = msg
		m.Quitting = true
		return m, tea.Quit

	}
	return m, cmd
}

func (m model) View() string {
	var str string
	if m.Err != nil {
		if errors.Is(m.Err, os.ErrExist) {
			str += common.Styles.Error.Render("Your .plan file does not exist; try running `plan init` to set up plan or `plan sync` to get your most recent plan from the server.")
		} else {
			str += common.Styles.Error.Render(m.Err.Error())
		}
	}
	if m.status == statusSuccess {
		str += "Done! Try running `plan sync` to synchronise your files with the server."
	}
	if m.Quitting {
		str += "\n"
	}
	return common.View.Render(str)
}
