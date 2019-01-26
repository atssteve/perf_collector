package collectors

import (
	"sync"

	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var allMeticCollectors = make(map[string]func() MetricsCollector)
var allConfigCollectors = make(map[string]func() ConfigCollector)
var activeMetricCollectors []string
var activeConfigCollectors []string

// MetricsCollector interface allows registeration of any collector simply by containing the UpdateMetrics receiver.
type MetricsCollector interface {
	UpdateMetrics(ch chan metrics.Metric)
}

// ConfigCollector interface allows registeration of any collector simply by containing the UpdateConfigs receiver.
type ConfigCollector interface {
	UpdateConfigs(ch chan metrics.Metric)
}

func registerAllMetricCollectors(collectorName string, collectorInit func() MetricsCollector) {
	allMeticCollectors[collectorName] = collectorInit
}

func registerAllConfigCollectors(collectorName string, collectorInit func() ConfigCollector) {
	allConfigCollectors[collectorName] = collectorInit
}

func findOutWhatsActive() {
	for collector, inbool := range viper.Get("collector.metric").(map[string]interface{}) {
		if inbool.(bool) {
			activeMetricCollectors = append(activeMetricCollectors, collector)
		}
	}
	for collector, inbool := range viper.Get("collector.config").(map[string]interface{}) {
		if inbool.(bool) {
			activeConfigCollectors = append(activeConfigCollectors, collector)
		}
	}
}

// LogActiveCollectors logs a list of registered collectors before kicking off the collection.
func LogActiveCollectors() {
	findOutWhatsActive()

	if len(activeMetricCollectors) > 0 {
		log.Infof("Registered Metric Collectors: %s", activeMetricCollectors)
	}
	if len(activeConfigCollectors) > 0 {
		log.Infof("Registered Config Collectors: %s", activeConfigCollectors)
	}

}

// UpdateMetricCollection requests all of the collectors to update their metrics.
func UpdateMetricCollection(ch chan metrics.Metric) {
	var wg sync.WaitGroup
	for _, col := range activeMetricCollectors {
		wg.Add(1)
		log.WithFields(log.Fields{
			"collector": col,
			"action":    "Starting Metric Collection",
		}).Info(col)
		collector := allMeticCollectors[col]()
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
	var wg sync.WaitGroup
	for _, col := range activeConfigCollectors {
		wg.Add(1)
		log.WithFields(log.Fields{
			"collector": col,
			"action":    "Starting Config Collection",
		}).Info(col)
		collector := allConfigCollectors[col]()
		go func() {
			collector.UpdateConfigs(ch)
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
}
