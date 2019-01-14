// +build darwin

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/cpu"
	log "github.com/sirupsen/logrus"
)

// GetCPUInfo collects all of the virtual memory information for the requested OS.
func GetCPUInfo(ch chan metrics.Metric) {
	v, err := cpu.Times(true)
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "cpu",
			"os":        "darwin",
			"action":    "GetCPUTimes",
		}).Errorf("Unable to get CPU stats: %+v", err)
	}

	log.WithFields(log.Fields{
		"collector": "cpu",
		"os":        "darwin",
	}).Debug(v)

	for _, metric := range v {
		cpu := metrics.CPU{
			CPU: metric,
		}
		ch <- cpu
	}
}
