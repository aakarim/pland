package fresh

import tea "github.com/charmbracelet/bubbletea"

func (m model) fresh() tea.Msg {
	if err := m.planService.Fresh(); err != nil {
		return err
	}
	return success{}
}
