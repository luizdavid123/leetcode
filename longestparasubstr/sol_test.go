package longestparasubstr

import (
	"leetcode/misc"
	"testing"
)

func TestLongestParaSubStr(t *testing.T) {
	tests := map[string]struct {
		in  string
		out string
	}{
		"case 01": {"babad", "bab"},
		"case 02": {"cbbd", "bb"},
		"case 03": {"a", "a"},
		"case 04": {"ac", "a"},
		"case 05": {"aa", "aa"},
		"case 06": {"bab", "bab"},
		"case 07": {"bccb", "bccb"},
		"case 08": {"abb", "bb"},
		"case 09": {"zasqweqwearanynarbasdasdas", "ranynar"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := LongestParaSubStr(test.in)
			misc.Equals(t, test.out, actual)
		})
	}
}

func TestIsStrPara(t *testing.T) {
	tests := map[string]struct {
		in  string
		out bool
	}{
		"case 01": {"a", true},
		"case 02": {"bb", true},
		"case 03": {"ab", false},
		"case 04": {"bab", true},
		"case 05": {"baab", true},
		"case 06": {"baaba", false},
		"case 07": {"aaaaa", true},
		"case 08": {"aaabaa", false},
		"case 09": {"aaabbaaa", true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := IsStrPara(test.in)
			misc.Equals(t, test.out, actual)
		})
	}
}
