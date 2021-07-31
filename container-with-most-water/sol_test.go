package sol

import (
	"leetcode/misc"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out int
	}{
		"case 01": {[]int{1, 1}, 1},
		"case 02": {[]int{4, 3, 2, 1, 4}, 16},
		"case 03": {[]int{1, 2, 1}, 2},
		"case 04": {[]int{1, 1}, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			act := MaxArea(test.in)
			misc.Equals(t, test.out, act)
		})
	}
}
