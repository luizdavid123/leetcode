package sol

import (
	"leetcode/misc"
	"testing"
)

func TestFindKthTop(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
		kmax int
	}{
		"case 01": {[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		"case 02": {[]int{3, 2, 3, 1, 2, 5, 4, 5, 6}, 4, 4},
		"case 03": {[]int{1}, 1, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			kmax := FindKthTopV2(test.nums, test.k)
			misc.Equals(t, test.kmax, kmax)
		})
	}
}
