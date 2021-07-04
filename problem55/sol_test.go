package problem55

import (
	"leetcode/misc"
	"testing"
)

func TestCanJump(t *testing.T) {
	tests := map[string]struct {
		nums []int
		cha  bool
	}{
		"case 01": {[]int{2, 3, 1, 1, 4}, true},
		"case 02": {[]int{2, 3, 1, 0, 4}, true},
		"case 03": {[]int{1}, true},
		"case 04": {[]int{2, 0}, true},
		"case 05": {[]int{1, 3, 0, 4}, true},
		"case 06": {[]int{2, 5, 0, 0}, true},
		"case 07": {[]int{2, 1, 0, 1}, false},
		"case 08": {[]int{3, 2, 1, 0, 4}, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cha := CanJumpV2(test.nums)
			misc.Equals(t, test.cha, cha)
		})
	}
}
