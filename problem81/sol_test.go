package problem81

import (
	"leetcode/misc"
	"testing"
)

func TestSearchInRotatedArr(t *testing.T) {
	tests := map[string]struct {
		nums  []int
		tar   int
		found bool
	}{
		"case 01": {[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		"case 02": {[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		"case 03": {[]int{1}, 1, true},
		"case 04": {[]int{1}, 0, false},
		"case 05": {[]int{1, 0, 1, 1, 1}, 0, true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			found := SearchInRotatedArr(test.nums, test.tar)
			misc.Equals(t, test.found, found)
		})
	}
}
