// +build linux

package collectors

import (
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

func GetMemInfo() {
	v, _ := mem.VirtualMemory()
	log.WithFields(log.Fields{
		"collector": "meminfo",
		"os":        "linux",
	}).Info(v)
}