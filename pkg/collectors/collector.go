package collectors

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

var registeredCollectors = make(map[string]func() Collector)

// Collector interface allows registeration of any collector simply by containing the Update receiver.
type Collector interface {
	Update(ch chan metrics.Metric)
}

func registerCollector(collectorName string, collectorInit func() Collector) {
	registeredCollectors[collectorName] = collectorInit
}

// StartCollection logs a list of registered collectors before kicking off the collection.
func StartCollection() {
	activeCollectors := []string{}
	for k := range registeredCollectors {
		activeCollectors = append(activeCollectors, k)
	}
	log.Infof("Registered Collectors: %s", activeCollectors)
}

// UpdateCollection requests all of the collectors to update their metrics.
func UpdateCollection(ch chan metrics.Metric) {
	for k, v := range registeredCollectors {
		log.WithFields(log.Fields{
			"collector": "meminfo",
			"action":    "Starting collection",
		}).Info(k)
		collector := v()
		collector.Update(ch)
	}
}
