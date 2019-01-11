// +build linux

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/disk"
	log "github.com/sirupsen/logrus"
)

// GetDiskInfo collects all of the available stats for disk
func GetDiskInfo(ch chan metrics.Metric) {
	filesystems, _ := disk.Partitions(false)

	for _, fsStats := range filesystems {
		v, _ := disk.Usage(fsStats.Mountpoint)
		log.WithFields(log.Fields{
			"collector": "disk",
			"os":        "linux",
		}).Info(v)

		ch <- v
	}
}
