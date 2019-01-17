package collectors

import (
	"time"

	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

var fsconf fsConfigData

func init() {
	registerCollector("filesystem", NewFSCollector)
	fsconf = fsConfigData{lastran: time.Now()}
}

type fsConfigData struct {
	lastran time.Time
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
func (m *fileSystemCollector) Update(ch chan metrics.Metric, interval *time.Duration) {
	switch {
	case time.Since(fsconf.lastran) < (time.Duration(1) * time.Second):
		log.WithFields(log.Fields{
			"collector":         "filesystem",
			"collection timing": interval,
			"action":            "Starting Collection",
		}).Info("filesystem")
		GetFSInfo(ch)
		fsconf.lastran = time.Now()

	case time.Since(fsconf.lastran).Seconds() <= interval.Seconds():
		return

	case time.Since(fsconf.lastran).Seconds() >= interval.Seconds():
		log.WithFields(log.Fields{
			"collector":         "filesystem",
			"collection timing": interval,
			"action":            "Continuing Collection",
		}).Info("filesystem")
		GetFSInfo(ch)
		fsconf.lastran = time.Now()
	}
}
