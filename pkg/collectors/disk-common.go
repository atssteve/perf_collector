package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

func init() {
	registerCollector("disk", NewDiskCollector)
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
func (m *diskCollector) Update(ch chan metrics.Metric) {
	log.WithFields(log.Fields{
		"collector": "disk",
		"action":    "Starting collection",
	}).Info("disk")
	GetDiskInfo(ch)

}
