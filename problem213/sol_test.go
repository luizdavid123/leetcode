package problem213

import (
	"leetcode/misc"
	"testing"
)

func TestRob(t *testing.T) {
	tests := map[string]struct {
		nums []int
		ans  int
	}{
		"case 01": {[]int{2, 3, 2}, 3},
		"case 02": {[]int{1, 2, 3, 1}, 4},
		"case 03": {[]int{0}, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := Rob(test.nums)
			misc.Equals(t, test.ans, ans)
		})
	}
}
