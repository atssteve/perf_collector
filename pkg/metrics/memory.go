package metrics

import (
	"encoding/json"

	"github.com/shirou/gopsutil/mem"
)

// Memory struct contains any memory information that needs to be collected.
type Memory struct {
	Memory *mem.VirtualMemoryStat `json:"memory"`
}

func (m Memory) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
