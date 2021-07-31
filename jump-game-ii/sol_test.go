package sol

import (
	"leetcode/misc"
	"testing"
)

func TestMinJump(t *testing.T) {
	tests := map[string]struct {
		nums []int
		num  int
	}{
		"case 02": {[]int{2, 3, 1, 1, 4}, 2},
		"case 03": {[]int{2, 3, 0, 1, 4}, 2},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			num := MinJumpV2(test.nums)
			misc.Equals(t, test.num, num)
		})
	}
}
