package sol

import (
	"leetcode/misc"
	"testing"
)

func TestCombSum(t *testing.T) {
	tests := map[string]struct {
		k    int
		n    int
		comb [][]int
	}{
		"case 01": {3, 7, [][]int{{1, 2, 4}}},
		"case 02": {3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		"case 03": {4, 1, [][]int{}},
		"case 04": {3, 2, [][]int{}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			comb := CombSum(test.k, test.n)
			misc.Equals(t, test.comb, comb)
		})
	}
}
