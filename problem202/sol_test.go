package sol

import (
	"leetcode/misc"
	"testing"
)

func TestIsHappy(t *testing.T) {
	tests := map[string]struct {
		num   int
		happy bool
	}{
		"case 01": {19, true},
		"case 02": {2, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			happy := IsHappy(test.num)
			misc.Equals(t, test.happy, happy)
		})
	}
}

func TestDigitSquereSum(t *testing.T) {
	tests := map[string]struct {
		n int
		s int
	}{
		"case 01": {1, 1},
		"case 02": {19, 82},
		"case 03": {22, 8},
		"case 04": {322, 17},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := DigitSquereSum(test.n)
			misc.Equals(t, test.s, s)
		})
	}
}
