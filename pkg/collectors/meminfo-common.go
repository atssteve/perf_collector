package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
)

func init() {
	registerMetricCollectors("meminfo", NewMemInfoCollector)
}

type memInfoCollector struct {
	MetricsCollector string
}

// NewMemInfoCollector creates a new memory collector for registration.
func NewMemInfoCollector() MetricsCollector {
	return &memInfoCollector{
		MetricsCollector: "meminfo",
	}
}

func (m *memInfoCollector) UpdateMetrics(ch chan metrics.Metric) {
	GetMemInfo(ch)
}
