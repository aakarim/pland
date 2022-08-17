/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/internal/plan"
	"github.com/aakarim/pland/cli/ui/sync"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync your plan files to the server.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		/**
			Sync requirements:
			- Runs whenever the file is changed, or when a notification is called. Possibly once a few mintes
			- needs to be debounced in case it is run lots in quick succession
			- Handles multi-machine conflicts using a branching 3-way-merge strategy
			Sync plan:
			1. Upload file to server with the parent id
				1. Server receives file
				2. Check from contents if this is a request for the latest entry or a historical import
					1. ...
				2. Get the parent version from the header
				3. Get the tip and check whether the tip id differs from the parent id
				4. if it does perform a merge with the tip of the latest version of that branch
					1. first check if hashes are the same, if they are then no merge necessary
					2. diff each section, if here are diffs store them in the file themselves for a user to resolve. These should contain the hash
					3. return the diff'd copy with a property of whether there is a conflict as a bool
				4. Assign the tip as the latest version with the parent as the one before it (it represents all the version on any branch in that time period)
			2. Receive most recent copy
			3. If there's a change overwite copy with copy from server
		**/
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

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
