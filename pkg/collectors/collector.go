package collectors

import (
	"fmt"
	"sync"
	"time"

	"github.com/atssteve/perf_collector/pkg/metrics"
)

// Currently this is a work in progress. Here we want to manage the intervals at which we collect metrics
// as well as any other possible configurations, sure as where to offload the data (file, network) and the
// most optimized way to orginize this data in a way to doesn't generate overhead on the end system but
// provides a flexible structure.

// CollectorConfig stores tunable configurations for the Linux Collection Process.
type CollectorConfig struct {
	Intervals time.Duration
}

// CollectorMetrics stores which metrics you wish to enable. This allows for pluggable metrics
// as well as disabling metrics that are unwanted.
type CollectorMetrics struct {
	MemInfo bool
}

// StartCollection starts up collection based on the details provided.
func StartCollection(c *CollectorConfig, m *CollectorMetrics) {
	fmt.Println("Starting collection on any node")
	var wg sync.WaitGroup
	time.Sleep(c.Intervals)
	switch {
	case m.MemInfo:
		wg.Add(1)
		go func() {
			defer wg.Done()
			getMetric("meminfo")
			getMetric("cpuinfo")

		}()
	}

	wg.Wait()
}

// getMetric will evalute the metric name passed and create a go routine and a buffered channel of bools
// which will wait for the go routine to complete its work.
func getMetric(m string) {
	done := make(chan bool, 1)
	switch m {
	case "meminfo":
		go metrics.GetMemInfo(done)
	case "cpuinfo":
		go metrics.CPUInfo(done)
	}
	<-done
}
