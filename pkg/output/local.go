package output

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/atssteve/perf_collector/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

const fileNamePrefix = "perf-collector-"

// Local contains info when writing to local filesystem
type Local struct {
	Enabled         bool
	Path            string
	RotationSize    string
	Compressed      bool
	RotationTime    time.Duration
	timeLastRotated time.Time
	workingFile     *os.File
	rw              sync.RWMutex
}

// Write drains the current channel to a file.
func (l *Local) Write(c chan metrics.Metric) {
	l.createFile()

	for {
		l.rw.Lock()
		l.rotateFile()
		m := <-c
		_, err := l.workingFile.WriteString(m.String() + "\n")
		if err != nil {
			panic(err)
		}
		l.rw.Unlock()
	}
}

// createFile either returns an existing file to use or creates a new one.
func (l *Local) createFile() {
	ts := time.Now()
	l.timeLastRotated = ts
	epoch := strconv.Itoa(int(ts.Unix()))
	filename := l.Path + "/" + fileNamePrefix + epoch
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		l.workingFile = file
	}
}

// rotateFile handles rotation of the file if needed.
func (l *Local) rotateFile() {
	if time.Since(l.timeLastRotated) > l.RotationTime {
		f, _ := l.workingFile.Stat()
		log.WithFields(log.Fields{
			"output": "local",
			"task":   "rotate",
		}).Infof("Rotating: %+v", f.Name())
		if l.Compressed {
			go l.compress(l.workingFile)
		} else {
			l.workingFile.Close()
		}
		l.createFile()
	}
}

// compress excepts a file to compress and cleans it up. Should be ran concurrently, not
// using the file from the struct to avoid having to add more locking.
func (l *Local) compress(file *os.File) {
	fz := file.Name() + ".gz"
	zip, err := os.Create(fz)
	if err != nil {
		log.WithFields(log.Fields{
			"output": "local",
			"task":   "compress",
			"action": "CreateZip",
		}).Errorf("Unable to open file: %+v", err)
	}
	defer os.Remove(file.Name())
	defer zip.Close()

	r := bufio.NewReader(file)
	content, err := ioutil.ReadAll(r)
	if err != nil {
		log.WithFields(log.Fields{
			"output": "local",
			"task":   "compress",
			"action": "OpenActiveFile",
		}).Errorf("Unable to read file: %+v", err)
	}
	w := gzip.NewWriter(zip)
	w.Write(content)
	w.Close()

}
