// +build linux

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/disk"
	log "github.com/sirupsen/logrus"
)

// GetDiskInfo collects all of the available stats for disk
func GetDiskInfo(ch chan metrics.Metric) {
	diskStats, _ := disk.IOCounters()
	log.WithFields(log.Fields{
		"collector": "disk",
		"os":        "linux",
	}).Info(diskStats)
	for _, v := range diskStats {
		ch <- v
	}
}
