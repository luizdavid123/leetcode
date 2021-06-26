package problem26

import (
	"leetcode/misc"
	"testing"
)

func TestRmDup(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out []int
		k   int
	}{
		"case 01": {[]int{1, 1, 2}, []int{1, 2, 1}, 2},
		"case 02": {[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4, 0, 2, 1, 3, 1}, 5},
		"case 03": {[]int{}, []int{}, 0},
		"case 04": {[]int{0}, []int{0}, 1},
		"case 05": {[]int{0, 1, 2}, []int{0, 1, 2}, 3},
		"case 06": {[]int{0, 1, 2, 2}, []int{0, 1, 2, 2}, 3},
		"case 07": {[]int{0}, []int{0}, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, k := RmDup(test.in)
			misc.Equals(t, test.k, k)
			misc.Equals(t, test.out, out)
		})
	}
}
