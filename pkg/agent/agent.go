package agent

import (
	"runtime"
	"time"

	"github.com/atssteve/perf_collector/pkg/collectors"
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/atssteve/perf_collector/pkg/output"
	log "github.com/sirupsen/logrus"
)

// This package manages the logistics of requesting updates to the collections based on
// what is passed in via cobra.

// Agent contains metadata about how the Agent has been requested to start.
type Agent struct {
	Intervals time.Duration
	Output    output.Output
	CPUTime   time.Duration
	DiskTime  time.Duration
	MemTime   time.Duration
	FSTime    time.Duration
}

var scheduler collectors.Scheduler

// Start is a prototype/placeholder right now.
func (a *Agent) Start() {
	// Were just passing down the times from the root agent to the scheduler
	scheduler.CPUTime = a.CPUTime
	scheduler.DiskTime = a.DiskTime
	scheduler.FSTime = a.FSTime
	scheduler.MemTime = a.MemTime

	// Making channels here for metrics and outputters
	go GetPerfData()
	localChan := make(chan metrics.Metric)

	log.WithFields(log.Fields{
		"pooling_intervals": a.Intervals,
	}).Info("Starting new agent")
	collectors.StartCollection()

	// Start up any enabled outputters
	if a.Output.Local.Enabled {
		go a.Output.Local.Write(localChan)
	}

	// Start collections
	// for {
	for x := 0; x < 12; x++ {
		metricsChannel := make(chan metrics.Metric, 1000)
		collectors.UpdateCollection(metricsChannel)
		for m := range metricsChannel {
			if a.Output.Local.Enabled {
				localChan <- m
			}
			// fmt.Println(m)
		}
		time.Sleep(1 * time.Second)

	}
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
