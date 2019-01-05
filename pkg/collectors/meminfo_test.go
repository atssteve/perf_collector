package collectors

import (
	"fmt"
	"testing"

	"github.com/atssteve/perf_collector/pkg/metrics"
)

// Test_parseMemInfo if make any changes still correctly parse the meminfo page. No point is making sure kB
// was parsed out correctly since that can't be an int64. Just using reflection to get the values of the map keys
// and making sure they will be usable. As changes to the /proc/meminfo output occurs, those edge cases should
// be tested here.
func Test_parseMemInfo(t *testing.T) {
	metricsChannel := make(chan metrics.Metric, 1)

	collector := NewMemInfoCollector()
	collector.Update(metricsChannel)
	result := <-metricsChannel
	fmt.Println(result.String())
}
