package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

func init() {
	registerCollector("cpu", NewCPUCollector)
}

type cpuCollector struct {
	Collector string
}

// NewCPUCollector creates a new cpu collector for registration.
func NewCPUCollector() Collector {
	return &cpuCollector{
		Collector: "cpu",
	}
}

//Update does things to cpu
func (m *cpuCollector) Update(ch chan metrics.Metric) {
	log.WithFields(log.Fields{
		"collector": "cpu",
		"action":    "Starting collection",
	}).Info("cpu")
	GetCPUInfo(ch)
}
