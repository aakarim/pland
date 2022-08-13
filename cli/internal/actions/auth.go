package actions

import (
	"github.com/aakarim/pland/cli/ui/auth"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/charm/client"
)

type NoUser struct{}
type UserFound struct {
	Name string
}
type AuthError struct {
	Err error
}

func (e AuthError) Error() string { return e.Err.Error() }
func (e AuthError) Unwrap() error { return e.Err }

func CheckUser() tea.Msg {
	c, err := client.NewClientWithDefaults()
	if err != nil {
		return AuthError{err}
	}
	u, err := c.Bio()
	if err != nil {
		return AuthError{err}
	}
	if u.Name == "" {
		return NoUser{}
	}
	return UserFound{
		Name: u.Name,
	}
}

func UpdateAuth(from tea.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	// return auth.InitialModel(from, nil), nil // TODO: remove
	switch msg := msg.(type) {
	case NoUser:
		return auth.InitialModel(from, nil), nil
	case AuthError:
		return auth.InitialModel(from, msg.Err), nil
	}
	return nil, nil
}
