package collectors

import (
	"sync"

	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup
var registeredMeticCollectors = make(map[string]func() MetricsCollector)
var registeredConfigCollectors = make(map[string]func() ConfigCollector)

// MetricsCollector interface allows registeration of any collector simply by containing the UpdateMetrics receiver.
type MetricsCollector interface {
	UpdateMetrics(ch chan metrics.Metric)
}

// ConfigCollector interface allows registeration of any collector simply by containing the UpdateConfigs receiver.
type ConfigCollector interface {
	UpdateConfigs(ch chan metrics.Metric)
}

func registerMetricCollectors(collectorName string, collectorInit func() MetricsCollector) {
	registeredMeticCollectors[collectorName] = collectorInit
}

func registerConfigCollectors(collectorName string, collectorInit func() ConfigCollector) {
	registeredConfigCollectors[collectorName] = collectorInit
}

// LogActiveCollectors logs a list of registered collectors before kicking off the collection.
func LogActiveCollectors() {
	activeMetricCollectors := []string{}
	activeConfigCollectors := []string{}
	for k := range registeredMeticCollectors {
		activeMetricCollectors = append(activeMetricCollectors, k)
	}
	log.Infof("Registered Metric Collectors: %s", activeMetricCollectors)

	for k := range registeredConfigCollectors {
		activeConfigCollectors = append(activeConfigCollectors, k)
	}
	log.Infof("Registered Config Collectors: %s", activeConfigCollectors)
}

// LogActiveConfigCollectors logs a list of registered collectors before kicking off the collection.
func LogActiveConfigCollectors() {
	activeConfigCollectors := []string{}
	for k := range registeredConfigCollectors {
		activeConfigCollectors = append(activeConfigCollectors, k)
	}
	log.Infof("Registered Config Collectors: %s", activeConfigCollectors)
}

// UpdateMetricCollection requests all of the collectors to update their metrics.
func UpdateMetricCollection(ch chan metrics.Metric) {
	for k, v := range registeredMeticCollectors {
		wg.Add(1)
		log.WithFields(log.Fields{
			"collector": k,
			"action":    "Starting Metric Collection",
		}).Info(k)
		collector := v()
		go func() {
			collector.UpdateMetrics(ch)
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
}

// UpdateConfigCollection requests all of the collectors to update their metrics.
func UpdateConfigCollection(ch chan metrics.Metric) {
	for k, v := range registeredConfigCollectors {
		wg.Add(1)
		log.WithFields(log.Fields{
			"collector": k,
			"action":    "Starting Config Collection",
		}).Info(k)
		collector := v()
		go func() {
			collector.UpdateConfigs(ch)
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
}
