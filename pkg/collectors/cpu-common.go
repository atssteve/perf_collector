package collectors

import (
	"time"

	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

var cpuconf cpuConfigData

func init() {
	registerCollector("cpu", NewCPUCollector)
	cpuconf = cpuConfigData{lastran: time.Now()}
}

type cpuConfigData struct {
	lastran time.Time
}

type cpuCollector struct {
	Collector string
}

// NewCPUCollector creates a new cpu collector for registration.
func NewCPUCollector() Collector {
	return &cpuCollector{
		Collector: "cpu",
	}
}

//Update does things to cpu
func (m *cpuCollector) Update(ch chan metrics.Metric, interval *time.Duration) {
	switch {
	case time.Since(cpuconf.lastran) < (time.Duration(1) * time.Second):
		log.WithFields(log.Fields{
			"collector":         "cpu",
			"collection timing": interval,
			"action":            "Starting Collection",
		}).Info("cpu")
		GetCPUInfo(ch)
		cpuconf.lastran = time.Now()

	case time.Since(cpuconf.lastran).Seconds() <= interval.Seconds():
		return

	case time.Since(cpuconf.lastran).Seconds() >= interval.Seconds():
		log.WithFields(log.Fields{
			"collector":         "cpu",
			"collection timing": interval,
			"action":            "Continuing Collection",
		}).Info("cpu")
		GetCPUInfo(ch)
		cpuconf.lastran = time.Now()
	}

}
