package problem167

import (
	"leetcode/misc"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct {
		nums []int
		tar  int
		ans  []int
	}{
		"case 01": {[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		"case 02": {[]int{2, 3, 4}, 6, []int{1, 3}},
		"case 03": {[]int{-1, 0}, -1, []int{1, 2}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := TwoSum(test.nums, test.tar)
			misc.Equals(t, test.ans, ans)
		})
	}
}
