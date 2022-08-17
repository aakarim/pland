/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/internal/plan"
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
		cfg := config.NewConfig(config.SetServer(config.ServerConfig{
			Host:        host,
			GraphQLPort: 8080,
			GraphQLPath: "/query",
			HttpScheme:  "http",
		}))
		ps := plan.NewPlanService(cfg)
		conflict, err := ps.Sync()
		if err != nil {
			return err
		}
		log.Println("synced; conflict:", conflict)
		return nil
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
