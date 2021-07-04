package problem53

import (
	"leetcode/misc"
	"testing"
)

func TestMaxSubSumV1(t *testing.T) {
	tests := map[string]struct {
		nums []int
		sum  int
	}{
		"case 01": {[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		"case 02": {[]int{1}, 1},
		"case 03": {[]int{5, 4, -1, 7, 8}, 23},
		"case 04": {[]int{-1}, -1},
		"case 05": {[]int{-1, -2}, -1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sum := MaxSubSumV1(test.nums)
			misc.Equals(t, test.sum, sum)
		})
	}
}
