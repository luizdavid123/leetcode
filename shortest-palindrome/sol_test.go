package sol

import (
	"leetcode/misc"
	"testing"
)

func TestShortestPalindrome(t *testing.T) {
	tests := map[string]struct {
		str string
		ans string
	}{
		"case 01": {"aacecaaa", "aaacecaaa"},
		"case 02": {"abcd", "dcbabcd"},
		"case 03": {"ab", "bab"},
		"case 04": {"bba", "abba"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := ShortestPalindrome(test.str)
			misc.Equals(t, test.ans, ans)
		})
	}
}

func TestRevStr(t *testing.T) {
	tests := map[string]struct {
		str string
		ans string
	}{
		"case 01": {"abc", "cba"},
		"case 02": {"abcd", "dcba"},
		"case 03": {"a", "a"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := RvsStr(test.str)
			misc.Equals(t, test.ans, ans)
		})
	}
}
