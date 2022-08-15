package auth

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/charm/client"
	charm "github.com/charmbracelet/charm/proto"
	charmCommon "github.com/charmbracelet/charm/ui/common"

	"github.com/aakarim/pland/cli/ui/common"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var viewStyle = lipgloss.NewStyle().Padding(1, 2, 2, 3)

type userUpdated struct{}

type model struct {
	common.Model
	from          tea.Model
	usernameInput textinput.Model
	updating      bool
	spinner       spinner.Model
}

func InitialModel(from tea.Model, err error) model {
	ti := textinput.New()
	ti.Placeholder = "Enter Username here"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	m := model{usernameInput: ti,
		spinner: charmCommon.NewSpinner(),
		from:    from,
		Model: common.Model{
			Err: err,
		}}
	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		spinner.Tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			m.Quitting = true
			return m, tea.Quit
		case tea.KeyEnter:
			m.Model.Err = nil
			if !m.updating {
				m.updating = true
				return m, m.setName
			}
			m.usernameInput.SetValue("")
			m.updating = false
			return m, nil
		}
	case error:
		m.Err = msg
		return m, nil
	case userUpdated:
		m.updating = false
		return m.from, nil
	default:
	}
	m.usernameInput, cmd = m.usernameInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	str := common.Keyword("Welcome to dotPlan! Create your account and start publishing your plans to the world.\n\n")
	if m.Err != nil {
		if errors.Is(m.Err, charm.ErrMissingSSHAuth) {
			return viewStyle.Render(str + common.Styles.Error.Render(
				"We were’t able to authenticate via SSH, which means there’s likely a problem with your key.\n\nYou can generate SSH keys by running "+common.Code("charm keygen")+". You can also set the environment variable "+common.Code("CHARM_SSH_KEY_PATH")+" to point to a specific private key, or use "+common.Code("-i")+"specifify a location."),
			)
		}
		return viewStyle.Render(
			str + "\n" +
				common.Styles.Error.Render(m.Err.Error()) +
				common.Styles.Paragraph.Render("\n\nPress enter to continue...\n"))
	}
	if m.updating {
		str += m.spinner.View()
	} else {
		str += fmt.Sprintf(
			common.Paragraph("What would you like your username to be?\n\n%s\n\n%s"),
			common.Paragraph(m.usernameInput.View()),
			"(esc to quit)")
	}
	if m.Quitting {
		return str + "\n"
	}
	return viewStyle.Render(str)

}

func (m model) setName() tea.Msg {
	c, err := client.NewClientWithDefaults()
	if err != nil {
		return err
	}
	_, err = c.SetName(m.usernameInput.Value())
	if err != nil {
		fmt.Println("returning err")
		return err
	}
	return userUpdated{}
}
