package problem6

import (
	"leetcode/misc"
	"testing"
)

func TestZigZagAndToRow(t *testing.T) {
	tests := map[string]struct {
		instr string
		inrow int
		out   string
	}{
		"case 01": {"a", 1, "a"},
		"case 02": {"a", 5, "a"},
		"case 03": {"ab", 1, "ab"},
		"case 04": {"abcd", 1, "abcd"},
		"case 05": {"abcd", 2, "acbd"},
		"case 06": {"abcd", 3, "abdc"},
		"case 07": {"abcd", 4, "abcd"},
		"case 08": {"paypalishiring", 4, "pinalsigyahrpi"},
		"case 09": {"paypalishiring", 3, "pahnaplsiigyir"},
		"case 10": {"abcdefg", 4, "agbfced"},
		"case 11": {"abcdefg", 3, "aebdfcg"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			act := ZigZagAndToRow(test.instr, test.inrow)
			misc.Equals(t, test.out, act)
		})
	}
}
