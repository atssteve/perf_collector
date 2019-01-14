// +build linux

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/disk"
	log "github.com/sirupsen/logrus"
)

// GetDiskInfo collects all of the available stats for disk
func GetDiskInfo(ch chan metrics.Metric) {
	diskStats, err := disk.IOCounters()
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "disk",
			"os":        "darwin",
			"action":    "GetDiskIO",
		}).Errorf("Unable to get disk stats: %+v", err)
	}
	log.WithFields(log.Fields{
		"collector": "disk",
		"os":        "darwin",
	}).Debug(diskStats)
	for _, v := range diskStats {
		disk := metrics.Disk{
			Disk: v,
		}
		ch <- disk
	}
}
