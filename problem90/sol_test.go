package problem90

import (
	"leetcode/misc"
	"testing"
)

func TestGenSubSetWithDup(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  [][]int
	}{
		"case 01": {[]int{0}, [][]int{{}, {0}}},
		"case 02": {[]int{0, 1}, [][]int{{}, {0}, {1}, {0, 1}}},
		"case 03": {[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		"case 04": {[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := GenSubSetWithDup(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
