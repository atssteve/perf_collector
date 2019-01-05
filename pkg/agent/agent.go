package agent

import (
	"fmt"
	"time"

	"github.com/atssteve/perf_collector/pkg/collectors"
	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

// This package manages the logistics of requesting updates to the collections based on
// what is passed in via cobra.

// Agent contains metadata about how the Agent has been requested to start.
type Agent struct {
	Intervals time.Duration
}

// Start is a prototype/placeholder right now.
func (a *Agent) Start() {
	metricsChannel := make(chan metrics.Metric, 1)
	log.WithFields(log.Fields{
		"pooling_intervals": a.Intervals,
	}).Info("Starting new agent")
	collectors.StartCollection()
	time.Sleep(a.Intervals)
	collectors.UpdateCollection(metricsChannel)
	fmt.Println(<-metricsChannel)

}
