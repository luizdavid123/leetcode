package sol

import (
	"leetcode/misc"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	tests := map[string]struct {
		in  int
		out string
	}{
		"case 01": {3, "III"},
		"case 02": {4, "IV"},
		"case 03": {9, "IX"},
		"case 04": {58, "LVIII"},
		"case 05": {1994, "MCMXCIV"},
		"case 06": {85, "LXXXV"},
		"case 07": {335, "CCCXXXV"},
		"case 08": {2222, "MMCCXXII"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			act := IntToRoman(test.in)
			misc.Equals(t, test.out, act)
		})
	}
}
