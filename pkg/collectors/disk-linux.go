// +build linux

package collectors

import (
	"encoding/json"

	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/disk"
	log "github.com/sirupsen/logrus"
)

//LoadDiskData Im al litte off the walls here
type LoadDiskData struct {
	DeviceName string
	FSInfo     *disk.UsageStat
	DevInfo    []disk.IOCountersStat
}

func (d LoadDiskData) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func alreadyInList(check string, devs []string) bool {
	for _, item := range devs {
		if item == check {
			return true
		}
	}
	return false
}
func GetDiskInfo(ch chan metrics.Metric) {
	filesystems, _ := disk.Partitions(true)
	devList := []string{}
	for _, fsStats := range filesystems {
		// devList = append(devList, fsStats.Device)
		if !alreadyInList(fsStats.Device, devList) {
			diskData := LoadDiskData{}
			diskData.DeviceName = fsStats.Device
			us, _ := disk.Usage(fsStats.Mountpoint)
			ios, _ := disk.IOCounters(fsStats.Device)

			diskData.FSInfo = us
			for _, diskDevData := range ios {
				diskData.DevInfo = append(diskData.DevInfo, diskDevData)
			}
			log.WithFields(log.Fields{
				"collector": "disk",
				"os":        "darwin",
			}).Info(diskData)

			ch <- diskData
		}
		devList = append(devList, fsStats.Device)

	}
}

// GetDiskInfo collects all of the available stats for disk
// func GetDiskInfo(ch chan metrics.Metric) {
// 	filesystems, _ := disk.Partitions(false)
// 	for _, fsStats := range filesystems {
// 		diskData := LoadDiskData{}
// 		diskData.DeviceName = fsStats.Device
// 		us, _ := disk.Usage(fsStats.Mountpoint)
// 		ios, _ := disk.IOCounters(fsStats.Device)

// 		diskData.FSInfo = us
// 		for _, diskDevData := range ios {
// 			diskData.DevInfo = append(diskData.DevInfo, diskDevData)
// 		}

// 		log.WithFields(log.Fields{
// 			"collector": "disk",
// 			"os":        "linux",
// 		}).Info(diskData)

// 		ch <- diskData
// 	}
// }
