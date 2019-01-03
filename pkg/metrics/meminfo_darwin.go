package metrics

import "fmt"

//GetMemInfo is testing the build process
func GetMemInfo(done chan bool) {
	fmt.Println("Will grab stats based file name for on Darwin Build")
	fmt.Println("Look at all my darwing stats........")
	done <- true
}
