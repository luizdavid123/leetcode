package sol

import (
	"leetcode/misc"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
		ans  []int
	}{
		"case 01": {[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		"case 02": {[]int{1}, 1, []int{1}},
		"case 03": {[]int{1, -1}, 1, []int{1, -1}},
		"case 04": {[]int{9, 11}, 2, []int{11}},
		"case 05": {[]int{4, -2}, 2, []int{4}},
		"case 06": {[]int{7, 2, 4}, 2, []int{7, 4}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := MaxSlidingWindow(test.nums, test.k)
			misc.Equals(t, test.ans, ans)
		})
	}
}
