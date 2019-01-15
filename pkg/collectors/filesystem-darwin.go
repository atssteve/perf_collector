// +build darwin

package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/disk"
	log "github.com/sirupsen/logrus"
)

// GetFSInfo collects all of the available stats for disk
func GetFSInfo(ch chan metrics.Metric) {
	mountedFS, err := disk.Partitions(false)
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "filesystem",
			"os":        "darwin",
			"action":    "GetFileSystems",
		}).Errorf("Unable to find mounted filesystems: %+v", err)
	}
	for _, FSs := range mountedFS {
		fsStats, err := disk.Usage(FSs.Mountpoint)
		if err != nil {
			log.WithFields(log.Fields{
				"collector": "filesystem",
				"os":        "darwin",
				"action":    "GetFSStats",
			}).Errorf("Unable to get stats from mounted filesystem: %+v", err)
		}
		log.WithFields(log.Fields{
			"collector": "filesystem",
			"os":        "darwin",
		}).Debug(fsStats)
		fsStat := metrics.FileSystem{
			FileSystem: *fsStats,
		}
		ch <- fsStat
	}
}
