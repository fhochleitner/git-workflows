package cmd

import (
	"gepaplexx/git-workflows/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of git-workflows",
	Long:  `All software has versions. This is git-workflows's`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("application version: %s", Version)
	},
}
