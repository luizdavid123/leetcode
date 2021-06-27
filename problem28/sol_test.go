package problem28

import (
	"leetcode/misc"
	"testing"
)

func TestStrIndxOf(t *testing.T) {
	tests := map[string]struct {
		str string
		sub string
		idx int
	}{
		"case 01": {"hello", "ll", 2},
		"case 02": {"aaaaa", "bba", -1},
		"case 03": {"", "", 0},
		"case 04": {"apple", "e", 4},
		"case 05": {"apple", "plee", -1},
		"case 06": {"", "ba", -1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := StrIndxOf(test.str, test.sub)
			misc.Equals(t, test.idx, actual)
		})
	}
}
