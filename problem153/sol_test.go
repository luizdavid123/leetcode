package sol

import (
	"leetcode/misc"
	"testing"
)

func TestFindMinInRotatedArr(t *testing.T) {
	tests := map[string]struct {
		nums []int
		min  int
	}{
		"case 01": {[]int{3, 4, 5, 6, 1, 2}, 1},
		"case 02": {[]int{4, 5, 6, 7, 8, 0, 1, 2}, 0},
		"case 03": {[]int{11, 13, 15, 17}, 11},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			min := FindMinInRotatedArr(test.nums)
			misc.Equals(t, test.min, min)
		})
	}
}
