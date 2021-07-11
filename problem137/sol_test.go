package problem137

import (
	"leetcode/misc"
	"testing"
)

func TestFindSingleNum(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  int
	}{
		"case 01": {[]int{2, 2, 3, 2}, 3},
		"case 02": {[]int{0, 1, 0, 1, 0, 1, 99}, 99},
		"case 03": {[]int{1}, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := FindSingleNum(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
