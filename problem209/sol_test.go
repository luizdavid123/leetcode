package problem209

import (
	"leetcode/misc"
	"testing"
)

func TestMinSubArrSumLen(t *testing.T) {
	tests := map[string]struct {
		nums []int
		tar  int
		len  int
	}{
		"case 01": {[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		"case 02": {[]int{1, 4, 4}, 4, 1},
		"case 03": {[]int{1, 1, 1, 1, 1, 1, 1}, 11, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			len := MinSubArrSumLen(test.nums, test.tar)
			misc.Equals(t, test.len, len)
		})
	}
}
