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

//Scheduler will hold the intervals each metric should collect for
type Scheduler struct {
	CPUTime  time.Duration
	DiskTime time.Duration
	MemTime  time.Duration
	FSTime   time.Duration
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

func schedule(f Collector, ch chan metrics.Metric, interval time.Duration, stop <-chan bool, wg *sync.WaitGroup) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		wg.Add(1)
		f.Update(ch)
		for {
			select {
			case <-ticker.C:
				f.Update(ch)

				wg.Done()
			case <-stop:
				return
			}
		}
	}()
	return ticker
}

// UpdateCollection requests all of the collectors to update their metrics.
func UpdateCollection(ch chan metrics.Metric, sch *Scheduler) {
	stop := make(chan bool)
	closeOut := map[string]*time.Ticker{}
	for k, v := range registeredCollectors {
		var timer time.Duration
		switch k {
		case "filesystem":
			// timer = sch.FSTime
			timer = time.Duration(6) * time.Second
		case "meminfo":
			// timer = sch.MemTime
			timer = time.Duration(2) * time.Second
		case "cpu":
			// timer = sch.CPUTime
			timer = time.Duration(8) * time.Second
		case "disk":
			// timer = sch.DiskTime
			timer = time.Duration(11) * time.Second
		}
		wg.Add(1)

		closeOut[k] = schedule(v(), ch, timer, stop, &wg)
	}
	fmt.Println(closeOut)
	wg.Wait()
	close(stop)
	for _, stillRunning := range closeOut {
		stillRunning.Stop()
	}
	close(ch)
}
