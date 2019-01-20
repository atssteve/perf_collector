package cmd

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/atssteve/perf_collector/pkg/agent"
	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

var rootCmd = &cobra.Command{
	Use:   "perf_collectory",
	Short: "Perf Collector is a plugable, portable metrics collector.",
	Long:  `A Plugable and Portable metrics collector.`,
	Run: func(cmd *cobra.Command, args []string) {
		newAgent.StartCollection()
	},
}

var newAgent agent.Agent

//Execute runs at the time the commandline tool is called.
func Execute() {
	rootCmd.Flags().BoolVar(&newAgent.Output.Local.Compressed, "no-compression", true, "Disable compression of rotated file")
	rootCmd.Flags().DurationVarP(&newAgent.MetricInterval, "metinterval", "i", time.Duration(2)*time.Second, "The number of seconds to wait before collecting metrics.")
	rootCmd.Flags().DurationVarP(&newAgent.ConfigInterval, "confinterval", "c", time.Duration(5)*time.Second, "The number of seconds to wait before collecting config.")
	rootCmd.Flags().BoolVar(&newAgent.Output.Local.Enabled, "local", false, "Output files to a local filesystems")
	rootCmd.Flags().StringVar(&newAgent.Output.Local.Path, "path", "/tmp", "Path in which to write files to.")
	rootCmd.Flags().DurationVar(&newAgent.Output.Local.RotationTime, "rotation-time", time.Duration(15)*time.Minute, "Duration in which to rotation file. Decreasing this entry will increase the rate which data is uploaded to S3")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
