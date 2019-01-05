package collectors

import (
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

func init() {
	registerCollector("meminfo", GetMemInfo)
}

// GetMemInfo will return a map containing the data parsed from /proc/meminfo
func GetMemInfo() {
	v, _ := mem.VirtualMemory()
	log.WithFields(log.Fields{
		"collector": "meminfo",
		"results":   "true",
	}).Info(v)
}
