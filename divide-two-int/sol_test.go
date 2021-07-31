package sol

import (
	"leetcode/misc"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := map[string]struct {
		dividend int
		divisor  int
		quotient int
	}{
		"case 01": {10, 3, 3},
		"case 02": {7, -3, -2},
		"case 03": {0, 1, 0},
		"case 04": {1, 1, 1},
		"case 05": {-8, 5, -1},
		"case 06": {-9, -4, 2},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Divide(test.dividend, test.divisor)
			misc.Equals(t, test.quotient, actual)
		})
	}
}
