package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
)

func init() {
	registerMetricCollectors("disk", NewDiskCollector)
}

type diskCollector struct {
	MetricsCollector string
}

// NewDiskCollector creates a new memory collector for registration.
func NewDiskCollector() MetricsCollector {
	return &diskCollector{
		MetricsCollector: "disk",
	}
}

func (m *diskCollector) UpdateMetrics(ch chan metrics.Metric) {
	GetDiskInfo(ch)
}
