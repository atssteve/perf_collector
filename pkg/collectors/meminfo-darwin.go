// +build darwin

package collectors

import (
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

// GetMemInfo collects all of the virtual memory information for the requested OS.
func GetMemInfo() {
	v, _ := mem.VirtualMemory()
	log.WithFields(log.Fields{
		"collector": "meminfo",
		"os":        "darwin",
	}).Info(v)
}
