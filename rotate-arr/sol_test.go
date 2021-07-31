package sol

import (
	"leetcode/misc"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		k      int
		rotate []int
	}{
		"case 01": {[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		"case 02": {[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		"case 03": {[]int{1, 2, 3, 4, 5, 6, 7}, 1, []int{7, 1, 2, 3, 4, 5, 6}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			rotate := Rotate(test.nums, test.k)
			misc.Equals(t, test.rotate, rotate)
		})
	}
}
