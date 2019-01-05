package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
)

func init() {
	registerCollector("cpu", NewCPUCollector)
}

type cpuCollector struct {
	Collector string
}

// NewCPUCollector creates a new memory collector for registration.
func NewCPUCollector() Collector {
	return &cpuCollector{
		Collector: "cpu",
	}
}

func (m *cpuCollector) Update(ch chan metrics.Metric) {
	GetCPUInfo(ch)
}
