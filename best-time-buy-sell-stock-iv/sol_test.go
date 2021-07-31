package sol

import (
	"leetcode/misc"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := map[string]struct {
		prices []int
		k      int
		max    int
	}{
		"case 01": {[]int{2, 4, 1}, 2, 2},
		"case 02": {[]int{3, 2, 6, 5, 0, 3}, 2, 7},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			max := MaxProfit(test.prices, test.k)
			misc.Equals(t, test.max, max)
		})
	}
}
