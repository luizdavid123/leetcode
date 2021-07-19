package problem152

import (
	"leetcode/misc"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := map[string]struct {
		nums []int
		max  int
	}{
		"case 01": {[]int{2, 3, -2, 4}, 6},
		"case 02": {[]int{-2, 0, -1}, 0},
		"case 03": {[]int{-2, 0, -1, 60}, 60},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			max := MaxProduct(test.nums)
			misc.Equals(t, test.max, max)
		})
	}
}
