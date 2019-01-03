package metrics

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// CPUTesting is for testing this cpu stuff
func CPUTesting() {
	t := time.Now()
	fmt.Println("This is a linux build")
	ctimes, _ := cpu.Times(false)
	mycpu, _ := cpu.Info()
	fmt.Println(mycpu)
	fmt.Println(ctimes)
	fmt.Println(time.Since(t))

}
