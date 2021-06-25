package problem20

import (
	"leetcode/misc"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := map[string]struct {
		input string
		valid bool
	}{
		"case 01": {"()", true},
		"case 02": {"()[]{}", true},
		"case 03": {"(]", false},
		"case 04": {"([)]", false},
		"case 05": {"{[]}", true},
		"case 06": {"(", false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ValidParentheses(test.input)
			misc.Equals(t, test.valid, actual)
		})
	}
}
