package problem7

import (
	"leetcode/misc"
	"testing"
)

func TestReverseNumber(t *testing.T) {
	tests := map[string]struct {
		in  int
		out int
	}{
		"case 01": {42, 24},
		"case 02": {123, 321},
		"case 03": {-42, -24},
		"case 04": {120, 21},
		"case 05": {0, 0},
		"case 06": {1200, 21},
		"case 07": {0, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			act := ReverseNumber(test.in)
			misc.Equals(t, test.out, act)
		})
	}
}
