package problem198

import (
	"leetcode/misc"
	"testing"
)

func TestMaxRob(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  int
	}{
		"case 01": {[]int{1, 2, 3, 1}, 4},
		"case 02": {[]int{2, 7, 9, 3, 1}, 12},
		"case 03": {[]int{2, 1, 1, 2}, 4},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := MaxRob(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
