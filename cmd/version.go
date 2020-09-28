package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display wnr version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(rootCmd.Use + " " + rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
