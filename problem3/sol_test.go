package problem3

import (
	"leetcode/misc"
	"testing"
)

func TestLongestSubStr(t *testing.T) {
	tests := map[string]struct {
		instr  string
		length int
	}{
		"case 01": {"abcabcbb", 3},
		"case 02": {"bbbbb", 1},
		"case 03": {"pwwkew", 3},
		"case 04": {"abcde", 5},
		"case 05": {"", 0},
		"case 06": {"aab", 2},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			len := LongestSubStr(test.instr)
			misc.Equals(t, test.length, len)
		})
	}
}
