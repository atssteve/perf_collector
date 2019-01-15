package collectors

import (
	"fmt"
	"sync"
	"time"

	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup
var registeredCollectors = make(map[string]func() Collector)
var fstimer time.Time

func init() {
	fmt.Println("HEYYYYYYY I RUN MORE THAN ONCE")
	fstimer = time.Now()
}

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
		// This is stop the filesystem stats from running every 2 seconds
		// right now the default is 1 minute. This is currently hard coded
		if k == "filesystem" {
			if time.Now().Sub(fstimer) > (time.Minute*1) || time.Now().Sub(fstimer) < (time.Second*2) {
				log.WithFields(log.Fields{
					"collector": k,
					"action":    "Starting collection",
				}).Info(k)
				collector := v()
				go func() {
					collector.Update(ch)
				}()
				fstimer = time.Now()
			}
		} else {
			wg.Add(1)
			log.WithFields(log.Fields{
				"collector": k,
				"action":    "Starting collection",
			}).Info(k)
			collector := v()
			go func() {
				collector.Update(ch)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	close(ch)
}
