package sol

import (
	"leetcode/misc"
	"testing"
)

func TestGenComb(t *testing.T) {
	tests := map[string]struct {
		n   int
		k   int
		ans [][]int
	}{
		"case 01": {4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		"case 02": {1, 1, [][]int{{1}}},
		"case 03": {4, 3, [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := GenCombV2(test.n, test.k)
			misc.Equals(t, test.ans, ans)
		})
	}
}
