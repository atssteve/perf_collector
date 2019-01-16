// +build darwin

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

// GetMemInfo collects all of the virtual memory information for the requested OS.
func GetMemInfo(ch chan metrics.Metric) {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "meminfo",
			"os":        "darwin",
			"action":    "GetMemory",
		}).Errorf("Unable to get memory stats: %+v", err)
	}
	memory := metrics.Memory{
		Memory: v,
	}
	log.WithFields(log.Fields{
		"collector": "meminfo",
		"os":        "darwin",
	}).Debug(memory)
	ch <- memory
}
