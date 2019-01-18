package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
)

func init() {
	registerMetricCollectors("disk", NewDiskCollector)
}

type diskCollector struct {
	Collector string
}

// NewDiskCollector creates a new memory collector for registration.
func NewDiskCollector() Collector {
	return &diskCollector{
		Collector: "disk",
	}
}

func (m *diskCollector) Update(ch chan metrics.Metric) {
	GetDiskInfo(ch)
}
