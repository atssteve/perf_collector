package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
)

func init() {
	registerConfigCollectors("disk", NewDiskCollector)
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
	GetFSInfo(ch)
}
