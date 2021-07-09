package problem72

import (
	"leetcode/misc"
	"testing"
)

func TestMinEditDist(t *testing.T) {
	tests := map[string]struct {
		lhs string
		rhs string
		ops int
	}{
		"case 01": {"horse", "ros", 3},
		"case 02": {"intention", "execution", 5},
		"case 03": {"", "", 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ops := MinEditDist(test.lhs, test.rhs)
			misc.Equals(t, test.ops, ops)
		})
	}
}
