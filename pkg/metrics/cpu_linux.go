package metrics

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// CPUInfo is for testing this cpu stuff
func CPUInfo(done chan bool) {
	t := time.Now()
	fmt.Println("Grabbing cpu stats for linux based on file name")
	ctimes, _ := cpu.Times(false)
	mycpu, _ := cpu.Info()
	fmt.Println(mycpu[0])
	fmt.Println(ctimes[0])
	fmt.Println(time.Since(t))
	done <- true

}
