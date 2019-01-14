package metrics

import (
	"encoding/json"

	"github.com/shirou/gopsutil/cpu"
)

// CPU struct contains any CPU information that needs to be collected.
type CPU struct {
	CPU cpu.TimesStat `json:"cpu"`
}

func (m CPU) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
