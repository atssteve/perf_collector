package output

import (
	"os"
	"strconv"
	"time"

	"github.com/atssteve/perf_collector/pkg/metrics"
)

const fileNamePrefix = "perf-collector-"

// Local contains info when writing to local filesystem
type Local struct {
	Enabled      bool
	Path         string
	RotationSize string
	Compressed   bool
}

// Write drains the current channel to a file.
func (l *Local) Write(c chan metrics.Metric) {
	file := l.createFile()
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for {
		m := <-c
		_, err := f.WriteString(m.String() + "\n")
		if err != nil {
			panic(err)
		}

	}
}

// createFile either returns an existing file to use or creates a new one.
func (l *Local) createFile() string {
	ts := time.Now().Unix()
	epoch := strconv.Itoa(int(ts))
	filename := l.Path + "/" + fileNamePrefix + epoch
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	return filename
}
