package cmd

import (
	"os"

	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/internal/plan"
	"github.com/aakarim/pland/cli/ui/sync"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var plainRender bool

var host = os.Getenv("CHARM_HOST")
var parentName = "plan"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "plan",
	Short: "Share what you're working on and get the world to engage",
	Long: `Create a new plan for today. 
	If you have an existing plan that was created yesterday, but edited today (perhaps you worked over midnight), it will save it as yesterday's file
	and wipe your current plan file for today.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var opts []tea.ProgramOption
		if plainRender {
			opts = append(opts, tea.WithoutRenderer())
		}
		cfg := config.NewConfig(config.SetServer(config.ServerConfig{
			Host:        host,
			GraphQLPort: 8080,
			GraphQLPath: "/query",
			HttpScheme:  "http",
		}))
		return tea.NewProgram(sync.InitialModel(plan.NewPlanService(cfg), cfg), opts...).Start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVarP(&plainRender, "plain", "d", false, "Disables fancy renderer and prints out results to the screen 'plainly'. Useful for grepping.")
}
