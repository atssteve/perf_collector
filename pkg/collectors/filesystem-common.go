package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
)

func init() {
	registerConfigCollectors("filesystem", NewFSCollector)
}

type fileSystemCollector struct {
	ConfigCollector string
}

// NewFSCollector creates a filesystem collector for registration.
func NewFSCollector() ConfigCollector {
	return &fileSystemCollector{
		ConfigCollector: "filesystem",
	}
}

//Update does things to fsstats
func (m *fileSystemCollector) UpdateConfigs(ch chan metrics.Metric) {
	GetFSInfo(ch)
}
