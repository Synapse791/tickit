package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/itmecho/tickit/types"
)

func init() {
	rootCmd.Flags().StringP("list-file", "l", UserHomeFile(".tickit.json"), "Location of the file that contains the list")
	viper.BindPFlag("list-file", rootCmd.Flags().Lookup("list-file"))
}

var taskList = types.TaskList{}

var rootCmd = &cobra.Command{
	Use:   "tickit",
	Short: "Command line todo list",
	PersistentPreRunE: func(cli *cobra.Command, args []string) error {
		if cli.Use == "init" {
			return nil
		}

		filePath := viper.GetString("list-file")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return fmt.Errorf("Please run the init command first")
		}

		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}

		return json.Unmarshal(fileData, &taskList)
	},
}

// Execute This is the main entrypoint to the cli
func Execute() error {
	return rootCmd.Execute()
}
