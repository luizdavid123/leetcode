package sol

import (
	"leetcode/misc"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := map[string]struct {
		s  string
		t  string
		is bool
	}{
		"case 01": {"anagram", "nagaram", true},
		"case 02": {"rat", "car", false},
		"case 03": {"a", "a", true},
		"case 04": {"a", "b", false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is := IsAnagram(test.s, test.t)
			misc.Equals(t, test.is, is)
		})
	}
}

func TestCntChFreq(t *testing.T) {
	tests := map[string]struct {
		str  string
		freq map[byte]int
	}{
		"case 01": {"foo", map[byte]int{'f': 1, 'o': 2}},
		"case 02": {"bar", map[byte]int{'b': 1, 'a': 1, 'r': 1}},
		"case 03": {"", map[byte]int{}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			freq := CntChFreq(test.str)
			misc.Equals(t, test.freq, freq)
		})
	}
}

func TestIsTwoMapEqual(t *testing.T) {
	tests := map[string]struct {
		m  map[byte]int
		n  map[byte]int
		is bool
	}{
		"case 01": {map[byte]int{}, map[byte]int{}, true},
		"case 02": {map[byte]int{'c': 1}, map[byte]int{'c': 1}, true},
		"case 03": {map[byte]int{'c': 1}, map[byte]int{'c': 0}, false},
		"case 04": {map[byte]int{'c': 1}, map[byte]int{'b': 0}, false},
		"case 05": {map[byte]int{'c': 1}, map[byte]int{'b': 1}, false},
		"case 06": {map[byte]int{'c': 1, 'b': 2}, map[byte]int{'b': 1}, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is := IsTwoMapEqual(test.m, test.n)
			misc.Equals(t, test.is, is)
		})
	}
}
