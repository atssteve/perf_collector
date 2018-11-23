package main

import (
	"fmt"
	"runtime"

	"github.com/atssteve/perf_collector/pkg/metrics"
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

	fmt.Println(metrics.GetMemInfo())

	runtime.ReadMemStats(&mem)
	log.WithFields(log.Fields{
		"Alloc":      mem.Alloc,
		"TotalAlloc": mem.TotalAlloc,
		"HeapAlloc":  mem.HeapAlloc,
		"HeapSys":    mem.HeapSys,
	}).Info("Ending Memory Statistics")
}
