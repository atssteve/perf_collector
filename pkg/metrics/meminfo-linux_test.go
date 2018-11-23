package metrics

import (
	"reflect"
	"strings"
	"testing"
)

// Test_parseMemInfo if make any changes still correctly parse the meminfo page. No point is making sure kB
// was parsed out correctly since that can't be an int64. Just using reflection to get the values of the map keys
// and making sure they will be usable. As changes to the /proc/meminfo output occurs, those edge cases should
// be tested here.
func Test_parseMemInfo(t *testing.T) {
	mockMemInfo := strings.NewReader("MemFree:\t1000 kB\nMem_Test:\t2000 kB\nActive(file):\t 3000 kB") // Example file testing key name parsing.
	memOut := parseMemInfo(mockMemInfo)
	for _, i := range *memOut {
		if len(i) < 1 {
			t.Error("Map has to many values. This shouldn't be possible.")
		}
		key := reflect.ValueOf(i).MapKeys()
		if strings.Contains(key[0].Interface().(string), "(") == true {
			t.Error("Found (, this should have been replaced via regex.")
		}
		if strings.Contains(key[0].Interface().(string), ")") == true {
			t.Error("Found ), this should have been replaced via regex.")
		}
	}
}
