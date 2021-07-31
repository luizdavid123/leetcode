package sol

import (
	"leetcode/misc"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := map[string]struct {
		n int
		a int
	}{
		"case 01": {2, 2},
		"case 02": {3, 3},
		"case 03": {1, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := ClimbStairsV2(test.n)
			misc.Equals(t, test.a, a)
		})
	}
}
