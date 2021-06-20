package paranumber

import (
	"leetcode/misc"
	"testing"
)

func TestIsPaliNum(t *testing.T) {
	tests := map[string]struct {
		in  int
		out bool
	}{
		"case 01": {42, false},
		"case 02": {123, false},
		"case 03": {-42, false},
		"case 04": {120, false},
		"case 05": {0, true},
		"case 06": {1200, false},
		"case 07": {121, true},
		"case 08": {11, true},
		"case 09": {1221, true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			act := IsPaliNum(test.in)
			misc.Equals(t, test.out, act)
		})
	}
}
