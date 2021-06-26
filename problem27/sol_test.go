package problem27

import (
	"leetcode/misc"
	"testing"
)

func TestRmDup(t *testing.T) {
	tests := map[string]struct {
		in  []int
		tar int
		out []int
		k   int
	}{
		"case 01": {[]int{1, 1, 2}, 1, []int{2, 1, 1}, 1},
		"case 02": {[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 2, []int{0, 0, 1, 1, 1, 3, 3, 4, 2, 2}, 8},
		"case 03": {[]int{}, 0, []int{}, 0},
		"case 04": {[]int{0}, 1, []int{0}, 1},
		"case 08": {[]int{0}, 0, []int{0}, 0},
		"case 05": {[]int{0, 1, 2}, 0, []int{2, 1, 0}, 2},
		"case 06": {[]int{0, 1, 2, 2}, 2, []int{0, 1, 2, 2}, 2},
		"case 07": {[]int{3, 2, 2, 3}, 3, []int{2, 2, 3, 3}, 2},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, k := RmEle(test.in, test.tar)
			misc.Equals(t, test.k, k)
			misc.Equals(t, test.out, out)
		})
	}
}
