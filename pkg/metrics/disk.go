package metrics

import (
	"encoding/json"

	"github.com/shirou/gopsutil/disk"
)

// Disk struct contains any Disk information that needs to be collected.
type Disk struct {
	Disk disk.IOCountersStat `json:"disk"`
}

func (m Disk) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
