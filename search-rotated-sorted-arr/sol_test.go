package sol

import (
	"leetcode/misc"
	"testing"
)

func TestFindIntInRotatedInts(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		idx    int
	}{
		"case 01": {[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		"case 02": {[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		"case 03": {[]int{1}, 0, -1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			idx := FindIntInRotatedInts(test.nums, test.target)
			misc.Equals(t, test.idx, idx)
		})
	}
}
