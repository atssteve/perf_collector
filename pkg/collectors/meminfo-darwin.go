// +build darwin

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

// GetMemInfo collects all of the virtual memory information for the requested OS.
func GetMemInfo() metrics.Metric {
	v, _ := mem.VirtualMemory()
	log.WithFields(log.Fields{
		"collector": "meminfo",
		"os":        "darwin",
	}).Info(v)

	return v
}
