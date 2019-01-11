// +build linux

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/cpu"
	log "github.com/sirupsen/logrus"
)

// GetMemInfo collects all of the virtual memory information for the requested OS.
func GetCPUInfo(ch chan metrics.Metric) {
	v, _ := cpu.Times(true)
	log.WithFields(log.Fields{
		"collector": "cpu",
		"os":        "linux",
	}).Info(v)

	for _, metric := range v {
		ch <- metric
	}
}
