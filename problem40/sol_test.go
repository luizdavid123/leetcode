package problem40

import (
	"leetcode/misc"
	"testing"
)

func TestCombSumUnique(t *testing.T) {
	tests := map[string]struct {
		nums []int
		tar  int
		ans  [][]int
	}{
		"case 01": {[]int{2, 3, 6, 7}, 7, [][]int{{7}}},
		"case 02": {[]int{2, 3, 5}, 8, [][]int{{3, 5}}},
		"case 03": {[]int{2}, 1, [][]int{}},
		"case 04": {[]int{1}, 1, [][]int{{1}}},
		"case 05": {[]int{1}, 2, [][]int{}},
		"case 06": {[]int{1, 2}, 3, [][]int{{1, 2}}},
		"case 07": {[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := CombSumUnique(test.nums, test.tar)
			misc.Equals(t, test.ans, ans)
		})
	}
}
