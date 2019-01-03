package metrics

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// CPUTesting is for testing this cpu stuff
func CPUTesting() {
	t := time.Now()
	fmt.Println("THis is the darwin build")
	mytimes, _ := cpu.Times(false)
	mycpu, _ := cpu.Info()
	fmt.Println(mycpu)
	fmt.Println(mytimes[0])
	fmt.Println(time.Since(t))

}
