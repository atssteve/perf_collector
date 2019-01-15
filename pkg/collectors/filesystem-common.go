package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
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

func (m *fileSystemCollector) Update(ch chan metrics.Metric) {
	GetFSInfo(ch)
}
