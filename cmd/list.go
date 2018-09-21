package cmd

import (
	"strings"

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

		for _, task := range taskList {
			status := " "
			if task.Complete {
				status = "x"
			}
			cli.Printf("[%s] - %s [%s]\n", status, task.Description, strings.Join(task.Tags, ","))
		}

		return nil
	},
}
