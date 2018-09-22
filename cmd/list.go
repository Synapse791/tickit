package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Prints the list of tasks",
	RunE: func(cli *cobra.Command, args []string) error {
		if len(taskList) == 0 {
			cli.Println("No tasks found")
			return nil
		}

		taskList.PrintTo(os.Stdout)

		return nil
	},
}
