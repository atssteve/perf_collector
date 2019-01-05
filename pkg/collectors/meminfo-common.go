package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
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

func (m *memInfoCollector) Update(ch chan metrics.Metric) {
	ch <- GetMemInfo()
}
