package sol

import (
	"leetcode/misc"
	"testing"
)

func TestFindSingleNum(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  int
	}{
		"case 01": {[]int{2, 2, 1}, 1},
		"case 02": {[]int{4, 1, 2, 1, 2}, 4},
		"case 03": {[]int{1}, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := FindSingleNum(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
