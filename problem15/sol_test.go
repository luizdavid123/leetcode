package problem15

import (
	"leetcode/misc"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out [][]int
	}{
		"case 01": {[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		"case 02": {[]int{}, [][]int{}},
		"case 03": {[]int{0}, [][]int{}},
		"case 04": {[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		"case 05": {[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ThreeSum(test.in, 0)
			misc.Equals(t, test.out, actual)
		})
	}
}

func TestIsAnsIncludesTripe(t *testing.T) {
	tests := map[string]struct {
		src [][]int
		tar []int
		out bool
	}{
		"case 04": {[][]int{{0, 0, 0}, {0, 0, 0}}, []int{0, 0, 0}, true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := IsAnsIncludesTripe(test.src, test.tar)
			misc.Equals(t, test.out, actual)
		})
	}
}
