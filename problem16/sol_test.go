package sol

import (
	"leetcode/misc"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		sum    int
	}{
		"case 01": {[]int{-1, 2, 1, -4}, 1, 2},
		"case 02": {[]int{0, 0, 0, 0}, 0, 0},
		"case 03": {[]int{0, 0, 0, 0}, 7, 0},
		"case 04": {[]int{0, 2, 1, -3}, 1, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ThreeSumClosest(test.nums, test.target)
			misc.Equals(t, test.sum, actual)
		})
	}
}
