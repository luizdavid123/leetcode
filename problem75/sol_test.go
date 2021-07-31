package sol

import (
	"leetcode/misc"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  []int
	}{
		"case 01": {[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		"case 02": {[]int{2, 0, 1}, []int{0, 1, 2}},
		"case 03": {[]int{1}, []int{1}},
		"case 04": {[]int{0}, []int{0}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := SortColors(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
