package problem34

import (
	"leetcode/misc"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		idxs   []int
	}{
		"case 01": {[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		"case 02": {[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		"case 03": {[]int{}, 6, []int{-1, -1}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			idxs := SearchRange(test.nums, test.target)
			misc.Equals(t, test.idxs, idxs)
		})
	}
}

func TestLowerBound(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		idx    int
	}{
		"case 01": {[]int{5, 7, 7, 8, 8, 10}, 8, 3},
		"case 02": {[]int{5, 7, 7, 8, 8, 10}, 7, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			idx := LowerBound(test.nums, test.target)
			misc.Equals(t, test.idx, idx)
		})
	}
}
