package cmd

import (
	"fmt"
	"os"
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
		/**
			// first load the existing plan file. If the plan file does not exist ask them to run `plan init`
			// update the file with a new entry
			// validate the existing plan file
				// there should be no > plan.header after a > plan.day
				// produces a warning if a plan contains
			// parse the plan file
				// find all the day locations
					// search through the file for all the `> plan.day` tokens with optional spaces
					// check the location to the next one as the ending bound (or EOF). Define those as start & end locations
					// find the first newline AFTER the token, that is the content start point
				// Find the header
					// locations 0 -> first `> plan.day`; OR
					// `> plan.header` -> first `> plan.day`
				// parse the contents of all the start and end locations and return
			// check most recent date of the plan file is a calendar day before in their timezone. If a person has flown to a timezone where they are the day before overnight they should be asked to take a nap!
			// update the parsed entity with a new entry
				// Add an empty day entry to the file
			// overwrite the existing file with the updated file
				// Create the string representation
					// create a buffer to be written
					// write the header
					// write the days
						// loop through the days writing each token as necessary
				// Write the bytes
			// recommend running 'sync' if the service isn't active
		**/
		var opts []tea.ProgramOption
		if plainRender {
			opts = append(opts, tea.WithoutRenderer())
		}
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("UserHomeDir(): %w", err)
		}
		cfg := config.NewConfig(config.SetServer(config.ServerConfig{
			Host:        os.Getenv("CHARM_HOST"),
			GraphQLPort: 8080,
			GraphQLPath: "/query",
			HttpScheme:  "http",
		}))
		cfg.ManagedPath = filepath.Join(homeDir, ".goplan/")
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
