package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.Flags().StringP("list-file", "l", "", "Location of the file that contains the list")
	viper.BindPFlag("list-file", rootCmd.Flags().Lookup("list-file"))
}

var rootCmd = &cobra.Command{
	Use:   "tickit",
	Short: "Command line todo list",
}

func Execute() error {
	return rootCmd.Execute()
}
