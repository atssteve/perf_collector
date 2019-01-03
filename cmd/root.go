package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/atssteve/perf_collector/pkg/collectors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "perf_collectory",
	Short: "Perf Collector is a plugable, portable metrics collector.",
	Long:  `A Plugable and Portable metrics collector.`,
	Run: func(cmd *cobra.Command, args []string) {
		collectors.StartCollection(&lcc, &lcm)
	},
}

var interval int64

var lcc = collectors.CollectorConfig{
	Intervals: time.Duration(interval),
}

var lcm = collectors.CollectorMetrics{
	MemInfo: true,
}

//Execute runs at the time the commandline tool is called.
func Execute() {
	lcc.Intervals = time.Duration(5) * time.Second
	lcm.MemInfo = true
	fmt.Println("Hey Im a common file")
	rootCmd.Flags().DurationVarP(&lcc.Intervals, "intervals", "i", time.Duration(5)*time.Second, "The number of seconds to wait before collecting metrics.")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
