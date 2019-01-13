package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/atssteve/perf_collector/pkg/agent"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "perf_collectory",
	Short: "Perf Collector is a plugable, portable metrics collector.",
	Long:  `A Plugable and Portable metrics collector.`,
	Run: func(cmd *cobra.Command, args []string) {
		newagent.Start()
	},
}

var newagent agent.Agent

//Execute runs at the time the commandline tool is called.
func Execute() {
	rootCmd.Flags().DurationVarP(&newagent.Intervals, "intervals", "i", time.Duration(2)*time.Second, "The number of seconds to wait before collecting metrics.")
	rootCmd.Flags().BoolVar(&newagent.Output.Local.Enabled, "local", false, "Output files to a local filesystems")
	rootCmd.Flags().StringVar(&newagent.Output.Local.Path, "path", "/tmp", "Path in which to write files to.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
