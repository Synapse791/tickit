package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(closeCmd)
}

var closeCmd = &cobra.Command{
	Use:   "close TASK_ID",
	Short: "Close a task as completed",
	Args:  cobra.ExactArgs(1),
	RunE: func(cli *cobra.Command, args []string) error {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		if err = taskList.CloseTask(taskID); err != nil {
			return err
		}

		if err = taskList.WriteToPath(viper.GetString("list-file")); err != nil {
			return err
		}

		cli.Printf("task with ID '%d' closed\n", taskID)
		return nil
	},
}
