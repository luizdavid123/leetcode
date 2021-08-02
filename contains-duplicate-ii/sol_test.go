package sol

import (
	"leetcode/misc"
	"testing"
)

func TestContainsNearbyDuplcate(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
		is   bool
	}{
		"case 01": {[]int{1, 2, 3, 1}, 3, true},
		"case 02": {[]int{1, 0, 1, 1}, 1, true},
		"case 03": {[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is := ContainsNearbyDuplcate(test.nums, test.k)
			misc.Equals(t, test.is, is)
		})
	}
}
