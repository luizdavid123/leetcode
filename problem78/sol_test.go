package problem78

import (
	"leetcode/misc"
	"testing"
)

func TestGenSubSet(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  [][]int
	}{
		// "case 01": {[]int{0}, [][]int{{}, {0}}},
		"case 02": {[]int{0, 1}, [][]int{{}, {0}, {1}, {0, 1}}},
		// "case 03": {[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := GenSubSet(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
