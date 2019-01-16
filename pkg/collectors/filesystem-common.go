package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

func init() {
	registerCollector("filesystem", NewFSCollector)
}

type fileSystemCollector struct {
	Collector string
}

// NewFSCollector creates a filesystem collector for registration.
func NewFSCollector() Collector {
	return &fileSystemCollector{
		Collector: "filesystem",
	}
}

//Update does things to fsstats
func (m *fileSystemCollector) Update(ch chan metrics.Metric) {
	log.WithFields(log.Fields{
		"collector": "filesystem",
		"action":    "Starting collection",
	}).Info("filesystem")
	GetFSInfo(ch)
}
