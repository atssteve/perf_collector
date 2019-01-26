package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
)

func init() {
	registerAllMetricCollectors("cpu", NewCPUCollector)
}

type cpuCollector struct {
	MetricsCollector string
}

// NewCPUCollector creates a new cpu collector for registration.
func NewCPUCollector() MetricsCollector {
	return &cpuCollector{
		MetricsCollector: "cpu",
	}
}

func (m *cpuCollector) UpdateMetrics(ch chan metrics.Metric) {
	GetCPUInfo(ch)
}
