package metrics

import (
	"encoding/json"

	"github.com/shirou/gopsutil/disk"
)

// FileSystem struct contains any information gathered about all non viotile
// ie /dev/shm /dev/run filesytems
type FileSystem struct {
	FileSystem disk.UsageStat `json:"filesystem"`
}

func (m FileSystem) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
