package main

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func main() {
	p, _ := process.Processes()
	for _, each := range p {
		fmt.Println(each.Times())
	}
}
