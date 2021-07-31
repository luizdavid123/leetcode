package sol

import (
	"leetcode/misc"
	"testing"
)

func TestCombSum(t *testing.T) {
	tests := map[string]struct {
		nums []int
		tar  int
		ans  [][]int
	}{
		"case 01": {[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		"case 02": {[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		"case 03": {[]int{2}, 1, [][]int{}},
		"case 04": {[]int{1}, 1, [][]int{{1}}},
		"case 05": {[]int{1}, 2, [][]int{{1, 1}}},
		"case 06": {[]int{1, 2}, 3, [][]int{{1, 1, 1}, {1, 2}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := CombSum(test.nums, test.tar)
			misc.Equals(t, test.ans, ans)
		})
	}
}
