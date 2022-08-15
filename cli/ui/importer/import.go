package importer

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/aakarim/pland/cli/internal/actions"
	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/internal/generated"
	"github.com/aakarim/pland/cli/internal/graphclient"
	"github.com/aakarim/pland/cli/ui/common"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/charm/client"
	"github.com/charmbracelet/lipgloss"
)

const (
	statusInit      = iota
	statusImporting = iota
	statusComplete  = iota
)

var viewStyle = lipgloss.NewStyle()

type importComplete struct {
	n int
}

type model struct {
	common.Model
	status    int
	path      string
	nImported int
}

func InitialModel(cfg config.Config, path string) model {
	return model{path: path, Model: common.Model{Config: cfg}}
}

func (m model) Init() tea.Cmd {
	return m.importImpl
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case error:
		return m, tea.Quit
	case importComplete:
		m.status = statusComplete
		m.nImported = msg.n
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	if m.Err != nil {
		return common.Styles.Error.Render(m.Err.Error()) + "\n"
	}
	str := common.Keyword("Welcome to pland! Create your account and start publishing your plans.\n\n")
	switch m.status {
	case statusInit:
		str += common.Paragraph("initialising...")
	case statusImporting:
		str += common.Paragraph("importing...")
	case statusComplete:
		str += common.Code(fmt.Sprintf("%d", m.nImported))
		str += common.Paragraph(" plans imported.")
	}

	return viewStyle.Render(str)
}

func (m model) importImpl() tea.Msg {
	var nImported int
	if err := filepath.WalkDir(m.path, func(path string, d fs.DirEntry, err error) error {
		ext := filepath.Ext(path)
		if ext == ".txt" || ext == ".md" {
			// reformat file and copy to dir
			fileName := strings.TrimSuffix(filepath.Base(path), ext)
			spl := strings.Split(fileName, "_")
			dateStr := spl[len(spl)-1]
			var foundKey bool
			// check if a valid if it's a plan file using the special prefix
			for _, v := range spl {
				if v == "plan" || v == "todo" {
					foundKey = true
				}
			}
			if !foundKey {
				fmt.Println("skipping non-plan file", path)
				return nil
			}
			// parse date
			t, err := time.Parse("20060102", dateStr)
			if err != nil {
				return fmt.Errorf("time.Parse(%s): %w", dateStr, err)
			}

			// get txt
			f, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("os.Open(%s): %w", path, err)
			}
			defer f.Close()
			b, err := ioutil.ReadAll(f)
			if err != nil {
				return fmt.Errorf("ioutil.ReadAll(%d): %w", len(b), err)
			}

			// publish to database
			c, err := client.NewClientWithDefaults()
			if err != nil {
				return actions.AuthError{Err: err}
			}
			j, err := c.JWT("charm")
			if err != nil {
				return err
			}
			u, err := url.Parse("http://localhost:8080/query")
			if err != nil {
				return err
			}
			httpClient := &http.Client{
				Transport: graphclient.NewAuthedTransport(j, http.DefaultTransport),
			}
			gqlClient := graphql.NewClient(u.String(), httpClient)
			resp, err := generated.CreatePlan(context.Background(), gqlClient, string(b), t)
			if err != nil {
				return fmt.Errorf("error creating: %w", err)
			}
			fmt.Println("imported", path, resp.CreatePlan.Id)
			nImported++
		}
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	fmt.Println("import complete")
	return importComplete{n: nImported}
}
