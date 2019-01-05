package collectors

import (
	log "github.com/sirupsen/logrus"
)

var registeredCollectors = make(map[string]func())

func registerCollector(collectorName string, collectorInit func()) {
	registeredCollectors[collectorName] = collectorInit
}

func StartCollection() {
	activeCollectors := []string{}
	for k := range registeredCollectors {
		activeCollectors = append(activeCollectors, k)
	}
	log.Infof("Registered Collectors: %s", activeCollectors)

	for k, v := range registeredCollectors {
		log.WithFields(log.Fields{
			"collector": "meminfo",
			"action":    "Starting collection",
		}).Info(k)
		v()
	}
}
