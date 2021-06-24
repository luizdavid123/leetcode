package problem18

import (
	"leetcode/misc"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		out    [][]int
	}{
		"case 02": {[]int{}, 0, [][]int{}},
		"case 01": {[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		"case 03": {[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := FourSum(test.nums, test.target)
			misc.Equals(t, test.out, actual)
		})
	}
}
