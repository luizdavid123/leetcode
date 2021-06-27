package problem35

import (
	"leetcode/misc"
	"testing"
)

func TestSearchOrInsertIfNotFound(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		idx    int
	}{
		"case 01": {[]int{1, 3, 5, 6}, 5, 2},
		"case 02": {[]int{1, 3, 5, 6}, 2, 1},
		"case 03": {[]int{1, 3, 5, 6}, 7, 4},
		"case 04": {[]int{1, 3, 5, 6}, 0, 0},
		"case 05": {[]int{1}, 0, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			idx := SearchOrInsertIfNotFound(test.nums, test.target)
			misc.Equals(t, test.idx, idx)
		})
	}
}
