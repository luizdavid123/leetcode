package sol

import (
	"leetcode/misc"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := map[string]struct {
		str  string
		dict []string
		ok   bool
	}{
		"case 01": {"leetcode", []string{"leet", "code"}, true},
		"case 02": {"applepenapple", []string{"apple", "pen"}, true},
		"case 03": {"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ok := WordBreak(test.str, test.dict)
			misc.Equals(t, test.ok, ok)
		})
	}
}
