package problem121

import (
	"leetcode/misc"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := map[string]struct {
		prices []int
		max    int
	}{
		"case 01": {[]int{7, 1, 5, 3, 6, 4}, 5},
		"case 02": {[]int{7, 6, 4, 3, 1}, 0},
		"case 03": {[]int{1}, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			max := MaxProfitV2(test.prices)
			misc.Equals(t, test.max, max)
		})
	}
}
