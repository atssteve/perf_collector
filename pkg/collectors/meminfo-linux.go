// +build linux

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

func GetMemInfo(ch chan metrics.Metric) {
	v, err := mem.VirtualMemory()
	log.WithFields(log.Fields{
		"collector": "meminfo",
		"os":        "linux",
		"action":    "GetMemory",
	}).Errorf("Unable to get memory stats: %+v", err)
	memory := metrics.Memory{
		Memory: v,
	}
	log.WithFields(log.Fields{
		"collector": "meminfo",
		"os":        "linux",
	}).Debug(memory)
	ch <- memory
}
