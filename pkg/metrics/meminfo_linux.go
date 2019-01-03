package metrics

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// GetMemInfo will return a map containing the data parsed from /proc/meminfo
func GetMemInfo(done chan bool) {
	fmt.Println("Grabbing Mem stats for Linux based on file name")
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "meminfo",
		}).Warn("Unable to access /proc/meminfo")
	}
	defer file.Close()

	log.WithFields(log.Fields{
		"collector": "meminfo",
		"results":   "true",
	}).Info(parseMemInfo(file))
	done <- true
}

// ParseMemInfo parses proc and returns a slice of maps containing all of the metrics keys and values.
func parseMemInfo(r io.Reader) *[]map[string]int64 {
	var mem []map[string]int64              // Create an empty slice to hold the metrics.
	var re = regexp.MustCompile(`\((.*)\)`) // Regex to replace any metrics wrapped in () ro be prefixed with _.

	fileScanner := bufio.NewScanner(r)
	for fileScanner.Scan() {
		m := make(map[string]int64)                   // Create a new map to hold the Metrics info.
		line := strings.Fields(fileScanner.Text())    // Split the line up, this will also remove the kB for us.
		key := line[0][:len(line[0])-1]               // Remove the trailing : from the metrics name.
		key = re.ReplaceAllString(key, "_$1")         // Changing any keys that included keys with ().
		value, _ := strconv.ParseInt(line[1], 10, 64) // Need to make sure value is on int64.
		m[key] = value
		mem = append(mem, m)
	}

	if fileScanner.Err() != nil {
		log.WithFields(log.Fields{
			"collector": "meminfo",
			"library":   "bufio",
			"function":  "scan",
		}).Warn("Unable read /proc/meminfo")
	}

	return &mem
}
