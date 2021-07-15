package problem123

import (
	"leetcode/misc"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := map[string]struct {
		prices []int
		max    int
	}{
		"case 01": {[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		"case 02": {[]int{1, 2, 3, 4, 5}, 4},
		"case 03": {[]int{7, 6, 4, 3, 1}, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			max := MaxProfit(test.prices)
			misc.Equals(t, test.max, max)
		})
	}
}
