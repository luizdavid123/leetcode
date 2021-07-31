package sol

import (
	"leetcode/misc"
	"testing"
)

func TestLongestSubStr(t *testing.T) {
	tests := map[string]struct {
		in  string
		out int
	}{
		"case 01": {"42", 42},
		"case 02": {"-42", -42},
		"case 03": {"   -42", -42},
		"case 04": {"4193 with word", 4193},
		"case 05": {"words and 987", 0},
		"case 06": {"-91283472332", -2147483648},
		"case 07": {"9223372036854775808", 2147483647},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			act := Atoi(test.in)
			misc.Equals(t, test.out, act)
		})
	}
}
