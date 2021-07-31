package sol

import (
	"leetcode/misc"
	"testing"
)

func TestFindMajorEle(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  int
	}{
		"case 01": {[]int{3, 2, 3}, 3},
		// "case 02": {[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := FindMajorEle(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
