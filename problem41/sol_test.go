package problem41

import (
	"leetcode/misc"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := map[string]struct {
		nums []int
		num  int
	}{
		"case 01": {[]int{1, 2, 0}, 3},
		"case 02": {[]int{3, 4, -1, 1}, 2},
		"case 03": {[]int{7, 8, 9, 11, 12}, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			num := FirstMissingPositive(test.nums)
			misc.Equals(t, test.num, num)
		})
	}
}
