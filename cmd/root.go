package cmd

import (
	"log"
	"time"

	"github.com/atssteve/perf_collector/config"
	"github.com/atssteve/perf_collector/pkg/agent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = cobra.Command{
	Use:   "perf_collector",
	Short: "Perf Collector is a plugable, portable metrics collector.",
	Long:  `A Plugable and Portable metrics collector.`,
	Run:   run,
}

var newAgent agent.Agent

// RootCommand is what allows the use of both flags and the config
// default order of priorty is flag, config, ENV, defaults
func RootCommand() *cobra.Command {
	rootCmd.Flags().BoolP("no-compression", "g", true, "Disable compression of rotated file")
	rootCmd.Flags().StringP("config", "c", "", "Override default config file. (default is .perf_config)")
	rootCmd.Flags().DurationP("intervals.metinterval", "i", time.Duration(2)*time.Second, "The number of seconds to wait before collecting metrics.")
	rootCmd.Flags().DurationP("intervals.confinterval", "o", time.Duration(5)*time.Second, "The number of seconds to wait before collecting config.")
	rootCmd.Flags().BoolP("local", "l", false, "Output files to a local filesystems")
	rootCmd.Flags().StringP("path", "p", "/tmp", "Path in which to write files to.")
	rootCmd.Flags().DurationP("rotation-time", "r", time.Duration(15)*time.Minute, "Duration in which to rotation file. Decreasing this entry will increase the rate which data is uploaded to S3")
	rootCmd.Flags().Bool("collector.metric.cpu", false, "Output files to a local filesystems")
	rootCmd.Flags().Bool("collector.metric.disk", false, "Output files to a local filesystems")
	rootCmd.Flags().Bool("collector.metric.mem", false, "Output files to a local filesystems")
	rootCmd.Flags().Bool("collector.config.filesystem", false, "Output files to a local filesystems")
	rootCmd.Flags().MarkHidden("collector.metric.cpu")
	rootCmd.Flags().MarkHidden("collector.metric.disk")
	rootCmd.Flags().MarkHidden("collector.metric.mem")
	rootCmd.Flags().MarkHidden("collector.config.filesystem")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	err := config.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	newAgent.MetricInterval = viper.GetDuration("intervals.metinterval")
	newAgent.ConfigInterval = viper.GetDuration("intervals.confinterval")

	newAgent.StartCollection()

}
