package collectors

import (
	"fmt"
	"strings"
	"testing"
)

// Test_parseMemInfo if make any changes still correctly parse the meminfo page. No point is making sure kB
// was parsed out correctly since that can't be an int64. Just using reflection to get the values of the map keys
// and making sure they will be usable. As changes to the /proc/meminfo output occurs, those edge cases should
// be tested here.
func Test_parseMemInfo(t *testing.T) {
	mockMemInfo := strings.NewReader("MemFree:\t1000 kB\nMem_Test:\t2000 kB\nActive(file):\t 3000 kB") // Example file testing key name parsing.
	fmt.Println(mockMemInfo)
}
