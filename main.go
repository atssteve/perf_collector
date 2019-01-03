package main

import (
	"github.com/atssteve/perf_collector/pkg/metrics"
)

func main() {
	metrics.CPUTesting()
	metrics.GetMemInfo()
}
