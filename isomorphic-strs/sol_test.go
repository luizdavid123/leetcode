package sol

import (
	"leetcode/misc"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	tests := map[string]struct {
		s  string
		t  string
		is bool
	}{
		"case 01": {"egg", "add", true},
		"case 02": {"foo", "bar", false},
		"case 03": {"paper", "title", true},
		"case 04": {"badc", "baba", false},
		"case 05": {"ab", "aa", false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is := IsIsomorphic(test.s, test.t)
			misc.Equals(t, test.is, is)
		})
	}
}
