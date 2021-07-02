package problem38

import (
	"leetcode/misc"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	tests := map[string]struct {
		n   int
		str string
	}{
		"case 01": {1, "1"},
		"case 02": {2, "11"},
		"case 03": {3, "21"},
		"case 04": {4, "1211"},
		"case 05": {5, "111221"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			str := CountAndSay(test.n)
			misc.Equals(t, test.str, str)
		})
	}
}
