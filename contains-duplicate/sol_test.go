package sol

import (
	"leetcode/misc"
	"testing"
)

func TestHasDup(t *testing.T) {
	tests := map[string]struct {
		nums []int
		is   bool
	}{
		"case 01": {[]int{1, 2, 3, 1}, true},
		"case 02": {[]int{1, 2, 3, 4}, true},
		"case 03": {[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is := HasDup(test.nums)
			misc.Equals(t, test.is, is)
		})
	}
}
