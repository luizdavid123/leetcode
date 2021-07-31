package sol

import (
	"leetcode/misc"
	"testing"
)

func TestAllPhoneNumToLetterComb(t *testing.T) {
	tests := map[string]struct {
		digits string
		combs  []string
	}{
		"case 01": {"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		"case 02": {"", []string{}},
		"case 03": {"2", []string{"a", "b", "c"}},
		"case 04": {"234", []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := AllPhoneNumToLetterComb(test.digits)
			misc.Equals(t, test.combs, actual)
		})
	}
}

func TestMerge(t *testing.T) {
	tests := map[string]struct {
		l []string
		r []string
		o []string
	}{
		"case 01": {[]string{"a", "b"}, []string{"d", "e"}, []string{"ad", "ae", "bd", "be"}},
		"case 02": {[]string{"a", "b", "c"}, []string{"d", "e", "f"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Merge(test.l, test.r)
			misc.Equals(t, test.o, actual)
		})
	}
}
