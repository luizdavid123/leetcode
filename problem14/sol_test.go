package sol

import (
	"leetcode/misc"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := map[string]struct {
		in  []string
		out string
	}{
		"case 01": {[]string{"flower", "flow", "flight"}, "fl"},
		"case 02": {[]string{"dog", "racecar", "car"}, ""},
		"case 03": {[]string{"flower", "flow", "flight", ""}, ""},
		"case 04": {[]string{"flower"}, "flower"},
		"case 05": {[]string{"dog", "dog", "dog"}, "dog"},
		"case 06": {[]string{"", ""}, ""},
		"case 07": {[]string{"flow", "flower", "flower"}, "flow"},
		"case 08": {[]string{""}, ""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			act := LongestCommonPrefix(test.in)
			misc.Equals(t, test.out, act)
		})
	}
}

func TestMaxLenStrInStrs(t *testing.T) {
	tests := map[string]struct {
		in  []string
		out string
	}{
		"case 01": {[]string{"flower", "flow", "flight"}, "flower"},
		"case 02": {[]string{"dog", "racecar", "car"}, "racecar"},
		"case 03": {[]string{"flower", "flow", "flight", ""}, "flower"},
		"case 04": {[]string{"flower"}, "flower"},
		"case 05": {[]string{"dog", "dog", "dog"}, "dog"},
		"case 06": {[]string{"", ""}, ""},
		"case 07": {[]string{"flow", "flower", "flower"}, "flower"},
		"case 08": {[]string{""}, ""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			act := MaxLenStrInStrs(test.in)
			misc.Equals(t, test.out, act)
		})
	}
}
