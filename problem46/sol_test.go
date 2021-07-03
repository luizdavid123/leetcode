package problem46

import (
	"leetcode/misc"
	"testing"
)

func TestPerm(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  [][]int
	}{
		"case 01": {[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		"case 02": {[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		"case 03": {[]int{1}, [][]int{{1}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := Perm(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
