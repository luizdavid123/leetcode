package problem32

import (
	"leetcode/misc"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := map[string]struct {
		str    string
		length int
	}{
		"case 01": {"(()", 2},
		"case 02": {")()())", 4},
		"case 03": {"", 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			length := LongestValidParentheses(test.str)
			misc.Equals(t, test.length, length)
		})
	}
}
