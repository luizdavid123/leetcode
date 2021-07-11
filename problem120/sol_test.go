package problem120

import (
	"leetcode/misc"
	"testing"
)

func TestMinTriPathSum(t *testing.T) {
	tests := map[string]struct {
		tri [][]int
		sum int
	}{
		"case 01": {[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		"case 02": {[][]int{{-10}}, -10},
		"case 03": {[][]int{{2}, {3, 4}}, 5},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sum := MinTriPathSum(test.tri)
			misc.Equals(t, test.sum, sum)
		})
	}
}
