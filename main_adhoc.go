package main

import (
	"runtime"

	"github.com/atssteve/perf_collector/pkg/collectors"
	log "github.com/sirupsen/logrus"
)

func main() {
	var mem runtime.MemStats

	runtime.ReadMemStats(&mem)
	log.WithFields(log.Fields{
		"Alloc":      mem.Alloc,
		"TotalAlloc": mem.TotalAlloc,
		"HeapAlloc":  mem.HeapAlloc,
		"HeapSys":    mem.HeapSys,
	}).Info("Starting Memory Statistics")

	lcc := &collectors.LinuxCollectorConfig{
		Intervals: 5,
	}

	lcm := &collectors.LinuxCollectorMetrics{
		MemInfo: true,
	}

	collectors.StartCollection(*lcc, *lcm)

	log.WithFields(log.Fields{
		"Alloc":      mem.Alloc,
		"TotalAlloc": mem.TotalAlloc,
		"HeapAlloc":  mem.HeapAlloc,
		"HeapSys":    mem.HeapSys,
	}).Info("Ending Memory Statistics")
}
