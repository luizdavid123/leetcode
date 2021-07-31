package sol

import (
	"leetcode/misc"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := map[string]struct {
		str string
		is  bool
	}{
		"case 01": {"A man, a plan, a canal: Panama", true},
		"case 02": {"race a car", false},
		"case 03": {"a", true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is := IsPalindrome(test.str)
			misc.Equals(t, test.is, is)
		})
	}
}

func TestIsDigitOrLetter(t *testing.T) {
	tests := map[string]struct {
		ch byte
		is bool
	}{
		"case 01": {'a', true},
		"case 02": {'A', true},
		"case 03": {'0', true},
		"case 04": {'l', true},
		"case 05": {'L', true},
		"case 06": {'8', true},
		"case 07": {':', false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			is := IsDigitOrLetter(test.ch)
			misc.Equals(t, test.is, is)
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := map[string]struct {
		str string
		out string
	}{
		"case 01": {"A man, a plan, a canal: Panama", "amanaplanacanalpanama"},
		"case 02": {"race a car", "raceacar"},
		"case 03": {"", ""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := Normalize(test.str)
			misc.Equals(t, test.out, out)
		})
	}
}
