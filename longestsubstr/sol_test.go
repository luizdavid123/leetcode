package longestsubstr

import (
	"leetcode/misc"
	"testing"
)

func TestLongestSubStr(t *testing.T) {
	tests := map[string]struct {
		instr  string
		outstr string
		length int
	}{
		"case 01": {"abcabcbb", "abc", 3},
		"case 02": {"bbbbb", "b", 1},
		"case 03": {"pwwkew", "wke", 3},
		"case 04": {"abcde", "abcde", 5},
		"case 05": {"", "", 0},
		"case 06": {"aab", "ab", 2},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			len, substr := LongestSubStr(test.instr)
			misc.Equals(t, test.outstr, substr)
			misc.Equals(t, test.length, len)
		})
	}
}
