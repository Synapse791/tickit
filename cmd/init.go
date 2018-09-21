package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Prints the list of tasks",
	RunE: func(cli *cobra.Command, args []string) error {
		filePath := viper.GetString("list-file")
		if _, err := os.Stat(filePath); err == nil {
			return fmt.Errorf("Already initialised")
		}

		cli.Println("creating list file...")
		return taskList.WriteToPath(viper.GetString("list-file"))
	},
}
