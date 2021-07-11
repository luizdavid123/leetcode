package problem118

import (
	"leetcode/misc"
	"testing"
)

func TestGenPascalTri(t *testing.T) {
	tests := map[string]struct {
		nRow int
		ans  [][]int
	}{
		"case 01": {1, [][]int{{1}}},
		"case 02": {2, [][]int{{1}, {1, 1}}},
		"case 03": {3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		"case 04": {4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := GenPascalTri(test.nRow)
			misc.Equals(t, test.ans, ans)
		})
	}
}
