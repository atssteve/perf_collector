package main

import (
	"github.com/atssteve/perf_collector/cmd"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	cmd.RootCommand().Execute()
}
