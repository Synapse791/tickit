package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(reopenCmd)
}

var reopenCmd = &cobra.Command{
	Use:   "reopen TASK_ID",
	Short: "Reopen a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cli *cobra.Command, args []string) error {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		if err = taskList.ReopenTask(taskID); err != nil {
			return err
		}

		if err = taskList.WriteToPath(viper.GetString("list-file")); err != nil {
			return err
		}

		cli.Printf("task %d reopened\n", taskID)
		return nil
	},
}
