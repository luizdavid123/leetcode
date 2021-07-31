package sol

import (
	"leetcode/misc"
	"testing"
)

func TestRomToInt(t *testing.T) {
	tests := map[string]struct {
		in  string
		out int
	}{
		"case 01": {"III", 3},
		"case 02": {"IV", 4},
		"case 03": {"IX", 9},
		"case 04": {"LVIII", 58},
		"case 05": {"MCMXCIV", 1994},
		"case 06": {"LXXXV", 85},
		"case 07": {"CCCXXXV", 335},
		"case 08": {"MMCCXXII", 2222},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			act := RomToInt(test.in)
			misc.Equals(t, test.out, act)
		})
	}
}
