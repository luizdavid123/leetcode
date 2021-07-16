package problem134

import (
	"leetcode/misc"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := map[string]struct {
		gas  []int
		cost []int
		idx  int
	}{
		"case 01": {[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		"case 02": {[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		"case 03": {[]int{5, 8, 2, 8}, []int{6, 5, 6, 6}, 3},
		"case 04": {[]int{5, 5, 1, 3, 4}, []int{8, 1, 7, 1, 1}, 3},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			idx := CanCompleteCircuit(test.gas, test.cost)
			misc.Equals(t, test.idx, idx)
		})
	}
}
