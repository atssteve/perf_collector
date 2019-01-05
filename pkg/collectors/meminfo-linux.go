// +build linux

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

func GetMemInfo(ch chan<- metrics.Metric) {
	v, _ := mem.VirtualMemory()
	log.WithFields(log.Fields{
		"collector": "meminfo",
		"os":        "linux",
	}).Info(v)
	ch <- v
}
