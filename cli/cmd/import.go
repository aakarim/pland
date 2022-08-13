package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aakarim/pland/cli/internal/config"
	"github.com/aakarim/pland/cli/ui/importer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:     "import",
	Short:   "import existing plan files from a folder to your managed database and publish",
	Long:    `Files must have the suffix plan_YYYYMMDD.txt or todo_YYYYMMDD.txt`,
	Args:    cobra.RangeArgs(0, 1),
	Example: indent.String(fmt.Sprintf("%s link\b%s link XXXXXX", parentName, parentName), 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var opts []tea.ProgramOption
		if plainRender {
			opts = append(opts, tea.WithoutRenderer())
		}
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("UserHomeDir(): %w", err)
		}

		cfg := config.Config{
			ManagedPath: filepath.Join(homeDir, ".goplan/"),
		}
		return tea.NewProgram(importer.InitialModel(cfg, args[0]), opts...).Start()
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
