package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

func init() {
	registerCollector("meminfo", NewMemInfoCollector)
}

type memInfoCollector struct {
	Collector string
}

// NewMemInfoCollector creates a new memory collector for registration.
func NewMemInfoCollector() Collector {
	return &memInfoCollector{
		Collector: "meminfo",
	}
}

//Update does things to mem
func (m *memInfoCollector) Update(ch chan metrics.Metric) {
	log.WithFields(log.Fields{
		"collector": "meminfo",
		"action":    "Starting collection",
	}).Info("meminfo")
	GetMemInfo(ch)
}
