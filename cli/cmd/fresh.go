package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/internal/plan"
	"github.com/aakarim/pland/cli/ui/fresh"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// freshCmd represents the fresh command
var freshCmd = &cobra.Command{
	Use:   "fresh",
	Short: "Start a fresh plan for the day",
	Long: `Start your day by running this command.

		It will add a new entry to your .plan or todo.txt file marked with > plan.day/YYYY-MM-DD // Pretty-printed time 

		The entry will be put at the top of the list of days.

		If you are on the same calendar day as the previous entry, it will do nothing.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var opts []tea.ProgramOption
		if plainRender {
			opts = append(opts, tea.WithoutRenderer())
		}
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("UserHomeDir(): %w", err)
		}
		var configFileOpt config.ConfigFunc
		f, err := os.Open(path.Join(homeDir, ".plan.cue"))
		if err == nil {
			defer f.Close()
			configFileOpt, err = config.WithFileValues(f)
			if err != nil {
				return err
			}
			f.Close()
		}
		cfg := config.NewConfig(configFileOpt, config.SetServer(config.ServerConfig{
			Host:        os.Getenv("CHARM_HOST"),
			GraphQLPort: 8080,
			GraphQLPath: "/query",
			HttpScheme:  "http",
		}), config.SetPlanPath(filepath.Join(homeDir, ".plan")))
		return tea.NewProgram(fresh.InitialModel(plan.NewPlanService(cfg), cfg), opts...).Start()
	},
}

func init() {
	rootCmd.AddCommand(freshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// freshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// freshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
