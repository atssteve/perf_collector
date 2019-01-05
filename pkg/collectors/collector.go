package collectors

import (
	log "github.com/sirupsen/logrus"
)

var registeredCollectors = make(map[string]func() Collector)

// Collector interface allows registeration of any collector simply by containing the Update receiver.
type Collector interface {
	Update()
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
	UpdateCollection()
}

// UpdateCollection requests all of the collectors to update their metrics.
func UpdateCollection() {
	for k, v := range registeredCollectors {
		log.WithFields(log.Fields{
			"collector": "meminfo",
			"action":    "Starting collection",
		}).Info(k)
		collector := v()
		collector.Update()
	}
}
