package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version number of Perf Collector",
	Long:  `Display the version number of Perf Collector`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Perf Collector version 0.1.")
	},
}
