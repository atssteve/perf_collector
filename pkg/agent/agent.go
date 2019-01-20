package agent

import (
	"runtime"
	"sync"
	"time"

	"github.com/atssteve/perf_collector/pkg/collectors"
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/atssteve/perf_collector/pkg/output"
	log "github.com/sirupsen/logrus"
)

// This package manages the logistics of requesting updates to the collections based on
// what is passed in via cobra.
var wg sync.WaitGroup

// Agent contains metadata about how the Agent has been requested to start.
type Agent struct {
	MetricInterval time.Duration
	ConfigInterval time.Duration
	Output         output.Output
}

// StartCollection kicks off all the collectors
func (a *Agent) StartCollection() {
	go GetPerfData()
	localChan := make(chan metrics.Metric)
	log.WithFields(log.Fields{
		"pooling_metric_interval": a.MetricInterval,
		"pooling_config_interval": a.ConfigInterval,
	}).Info("Starting new agent")
	collectors.LogActiveCollectors()

	// Start up any enabled outputters
	if a.Output.Local.Enabled {
		go a.Output.Local.Write(localChan)
	}

	go func() {
		for {
			metricsChannel := make(chan metrics.Metric, 1000)
			collectors.UpdateMetricCollection(metricsChannel)
			for m := range metricsChannel {
				if a.Output.Local.Enabled {
					localChan <- m
				}
			}
			time.Sleep(a.MetricInterval)
		}
	}()
	go func() {
		for {
			configChannel := make(chan metrics.Metric, 1000)
			collectors.UpdateConfigCollection(configChannel)
			for m := range configChannel {
				if a.Output.Local.Enabled {
					localChan <- m
				}
			}
			time.Sleep(a.ConfigInterval)
		}
	}()
	wg.Add(1)
	wg.Wait()
}

// GetPerfData logs current memory usage.
func GetPerfData() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		log.WithFields(log.Fields{
			"memory_alloc":   bToMB(m.Alloc),
			"memory_total":   bToMB(m.TotalAlloc),
			"memory_heap":    bToMB(m.HeapAlloc),
			"memory_objects": m.HeapObjects,
			"memory_sys":     bToMB(m.Sys),
			"memory_num_gc":  m.NumGC,
		}).Info("Memory Statistics")
		time.Sleep(time.Second * 5)
	}
}

func bToMB(b uint64) uint64 {
	return b / 1024 / 1024
}
