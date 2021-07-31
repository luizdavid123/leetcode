package sol

import (
	"leetcode/misc"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  string
	}{
		"case 01": {[]int{10, 2}, "210"},
		"case 02": {[]int{1}, "1"},
		"case 03": {[]int{10}, "10"},
		"case 04": {[]int{3, 30, 34, 5, 9}, "9534330"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := LargestNumber(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
