package sol

import (
	"leetcode/misc"
	"testing"
)

func TestKeepUniqueNumAtMostTwo(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
	}{
		"case 01": {[]int{1, 1, 1, 2, 2, 3}, 5},
		"case 02": {[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
		"case 03": {[]int{1}, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			k := KeepUniqueNumAtMostTwo(test.nums)
			misc.Equals(t, test.k, k)
		})
	}
}
