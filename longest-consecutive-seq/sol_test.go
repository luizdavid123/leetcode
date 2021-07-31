package sol

import (
	"leetcode/misc"
	"testing"
)

func TestLongestConsective(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  int
	}{
		"case 01": {[]int{100, 4, 200, 1, 3, 2}, 4},
		"case 02": {[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		"case 03": {[]int{0}, 1},
		"case 04": {[]int{}, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := LongestConsective(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
