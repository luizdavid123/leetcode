package sol

import (
	"leetcode/misc"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  []int
	}{
		"case 01": {[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		"case 02": {[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := ProductExceptSelf(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
