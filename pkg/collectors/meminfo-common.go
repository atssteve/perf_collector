package collectors

import (
	"time"

	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

var memconf memConfigData

func init() {
	registerCollector("meminfo", NewMemInfoCollector)
	memconf = memConfigData{lastran: time.Now()}
}

type memConfigData struct {
	lastran time.Time
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

//Update does things to mem
func (m *memInfoCollector) Update(ch chan metrics.Metric, interval *time.Duration) {
	switch {
	case time.Since(memconf.lastran) <= (time.Duration(1) * time.Second):
		log.WithFields(log.Fields{
			"collector":         "meminfo",
			"collection timing": interval,
			"action":            "Starting Collection",
		}).Info("meminfo")
		GetMemInfo(ch)
		memconf.lastran = time.Now()

	case time.Since(memconf.lastran).Seconds() <= interval.Seconds():
		return
	case time.Since(memconf.lastran).Seconds() >= interval.Seconds():
		log.WithFields(log.Fields{
			"collector":         "meminfo",
			"collection timing": interval,
			"action":            "Continuing Collection",
		}).Info("meminfo")
		GetMemInfo(ch)
		memconf.lastran = time.Now()
	}
}
