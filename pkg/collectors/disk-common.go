package collectors

import (
	"time"

	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

var diskconf diskConfigData

func init() {
	registerCollector("disk", NewDiskCollector)
	diskconf = diskConfigData{lastran: time.Now()}
}

type diskConfigData struct {
	lastran time.Time
}
type diskCollector struct {
	Collector string
}

// NewDiskCollector creates a disk collector for registration.
func NewDiskCollector() Collector {
	return &diskCollector{
		Collector: "disk",
	}
}

//Update does things to disk
func (m *diskCollector) Update(ch chan metrics.Metric, interval *time.Duration) {
	switch {
	case time.Since(fsconf.lastran) < (time.Duration(1) * time.Second):
		log.WithFields(log.Fields{
			"collector":         "disk",
			"collection timing": interval,
			"action":            "Starting Collection",
		}).Info("disk")
		GetDiskInfo(ch)
		diskconf.lastran = time.Now()

	case time.Since(fsconf.lastran).Seconds() <= interval.Seconds():
		return

	case time.Since(diskconf.lastran).Seconds() >= interval.Seconds():
		log.WithFields(log.Fields{
			"collector":         "disk",
			"collection timing": interval,
			"action":            "Continuing Collection",
		}).Info("disk")
		GetDiskInfo(ch)
		diskconf.lastran = time.Now()
	}
}
