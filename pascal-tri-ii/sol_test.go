package sol

import (
	"leetcode/misc"
	"testing"
)

func TestGetRowOfPascalTri(t *testing.T) {
	tests := map[string]struct {
		row int
		ans []int
	}{
		"case 01": {0, []int{1}},
		"case 02": {1, []int{1, 1}},
		"case 03": {2, []int{1, 2, 1}},
		"case 04": {3, []int{1, 3, 3, 1}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := GetRowOfPascalTri(test.row)
			misc.Equals(t, test.ans, ans)
		})
	}
}
