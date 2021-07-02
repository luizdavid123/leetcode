package problem49

import (
	"leetcode/misc"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := map[string]struct {
		strs []string
		ans  [][]string
	}{
		"case 01": {[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		"case 02": {[]string{""}, [][]string{{""}}},
		"case 03": {[]string{"a"}, [][]string{{"a"}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := GroupAnagrams(test.strs)
			misc.Equals(t, test.ans, ans)
		})
	}
}

func TestIsTwoStrAnag(t *testing.T) {
	tests := map[string]struct {
		lhs string
		rhs string
		is  bool
	}{
		"case 01": {"a", "a", true},
		"case 02": {"tea", "ate", true},
		"case 03": {"", "", true},
		"case 04": {"", "tea", false},
		"case 05": {"bar", "foo", false},
		"case 06": {"flow", "fl", false},
		"case 07": {"f", "flow", false},
		"case 08": {"tan", "nat", true},
		"case 09": {"dddg", "gggd", false},
		"case 10": {"duh", "ill", false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is := IsTwoStrAnag(test.lhs, test.rhs)
			misc.Equals(t, test.is, is)
		})
	}
}
