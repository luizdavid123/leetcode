package problem150

import (
	"leetcode/misc"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := map[string]struct {
		tokens []string
		value  int
	}{
		"case 01": {[]string{"2", "1", "+", "3", "*"}, 9},
		"case 02": {[]string{"4", "13", "5", "/", "+"}, 6},
		"case 03": {[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := EvalRPN(test.tokens)
			misc.Equals(t, test.value, value)
		})
	}
}
