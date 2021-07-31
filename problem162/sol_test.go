package sol

import (
	"leetcode/misc"
	"testing"
)

func TestFindPeakEle(t *testing.T) {
	tests := map[string]struct {
		nums []int
		peak int
	}{
		"case 01": {[]int{1, 2, 3, 1}, 2},
		"case 02": {[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		"case 03": {[]int{2}, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			peak := FindPeakEle(test.nums)
			misc.Equals(t, test.peak, peak)
		})
	}
}
