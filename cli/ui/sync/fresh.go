package sync

import (
	"errors"

	"github.com/aakarim/pland/cli/internal/actions"
	"github.com/aakarim/pland/cli/internal/plan"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/charm/proto"
)

func (m model) fresh() tea.Msg {
	status, err := m.planService.CopyOrCreatePlanFromStore(m.planPath)
	if err != nil {
		if errors.Is(err, proto.ErrMissingSSHAuth) || errors.Is(err, proto.ErrMissingUser) || errors.Is(err, proto.ErrMalformedKey) {
			return actions.AuthError{Err: err}
		}
	}
	var s int
	switch status {
	case plan.StatusFinishedWithClean:
		s = statusFinishedWithClean
	case plan.StatusFinishedWithCopy:
		s = statusFinishedWithCopy
	}
	return freshPlanCreated{path: m.planPath, status: s}
}
