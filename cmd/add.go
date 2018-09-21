package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/itmecho/tickit/types"
)

func init() {
	addCmd.Flags().StringSliceP("tags", "t", []string{}, "Tags to add to the task")
	viper.BindPFlag("add-tags", addCmd.Flags().Lookup("tags"))

	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [options] TASK",
	Short: "Add a task to the list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cli *cobra.Command, args []string) error {
		description := args[0]
		tags := viper.GetStringSlice("add-tags")

		newTask := &types.Task{
			Description: description,
			Complete:    false,
			Tags:        tags,
		}

		taskList = append(taskList, newTask)
		if err := taskList.WriteToPath(viper.GetString("list-file")); err != nil {
			return err
		}

		cli.Printf("added task \"%s\" with %d tags (%s)\n", description, len(tags), strings.Join(tags, ","))
		return nil
	},
}
