// +build darwin

package collectors

import (
	"encoding/json"
	"fmt"

	"github.com/atssteve/perf_collector/pkg/metrics"
	"github.com/shirou/gopsutil/disk"
	log "github.com/sirupsen/logrus"
)

//LoadDiskData Im al litte off the walls here
type LoadDiskData struct {
	DeviceName string
	FSInfo     disk.UsageStat
}

func (d LoadDiskData) String() string {
	// myret := make(map[string]disk.UsageStat)
	s, _ := json.Marshal(d)
	return string(s)
}

// GetDiskInfo collects all of the available stats for disk
func GetDiskInfo(ch chan metrics.Metric) {
	filesystems, _ := disk.Partitions(false)
	for _, fsStats := range filesystems {
		myNewGuy := LoadDiskData{}
		myNewGuy.DeviceName = fsStats.Device
		fmt.Println("HELLO", fsStats)
		v, _ := disk.Usage(fsStats.Mountpoint)
		myNewGuy.FSInfo = *v
		log.WithFields(log.Fields{
			"collector": "disk",
			"os":        "darwin",
		}).Info(myNewGuy)
		// customData := make(map[string]disk.UsageStat)
		// customData[fsStats.Device] = *v
		ch <- myNewGuy
		// ch <- v
	}
}
