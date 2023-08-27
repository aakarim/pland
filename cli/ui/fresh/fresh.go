package fresh

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) fresh() tea.Msg {
	if err := m.planService.Fresh(); err != nil {
		return err
	}
	return success{}
}

func (m model) openEditor() tea.Msg {
	if m.cfg.EditorCommand != "" {
		p, err := exec.LookPath(m.cfg.EditorCommand)
		if err != nil {
			return fmt.Errorf("opening editor: %w", err)
		}
		cmd := exec.Command(p, m.cfg.PlanPath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("running editor: %w", err)
		}
		return openedEditor{}
	}
	return didNotOpenEditor{}
}
