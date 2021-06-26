package problem22

import (
	"leetcode/misc"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	tests := map[string]struct {
		in  int
		out []string
	}{
		// "case 01": {1, []string{"()"}},
		"case 02": {2, []string{"()()", "(())"}},
		// "case 03": {3, []string{"()()()", "(())()", "(()())", "()(())", "((()))"}},
		// "case 04": {4, GenerateValidParenthesesV2(4)},
		// "case 05": {5, GenerateValidParenthesesV2(5)},
		// "case 06": {6, GenerateValidParenthesesV2(6)},
		// "case 08": {8, GenerateValidParenthesesV2(8)},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := GenerateValidParenthesesV3(test.in)
			misc.Equals(t, true, ValidAns(test.out, actual))
		})
	}
}

// ValidAns check if the ans is valid
func ValidAns(expect, actual []string) bool {
	if len(expect) != len(actual) {
		return false
	}
	for i := 0; i < len(actual); i++ {
		if !IsStrInStrs(expect, actual[i]) {
			return false
		}
	}
	return true
}

// IsStrInStrs check if the str is in strs
func IsStrInStrs(strs []string, str string) bool {
	for i := 0; i < len(strs); i++ {
		if strs[i] == str {
			return true
		}
	}
	return false
}
