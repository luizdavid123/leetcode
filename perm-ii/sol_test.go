package sol

import (
	"leetcode/misc"
	"testing"
)

func TestUniquePerm(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  [][]int
	}{
		"case 01": {[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		"case 02": {[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := UniquePerm(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
