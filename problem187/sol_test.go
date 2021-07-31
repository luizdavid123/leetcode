package sol

import (
	"leetcode/misc"
	"testing"
)

func TestFindRepeatedDNASeq(t *testing.T) {
	tests := map[string]struct {
		str string
		ans []string
	}{
		"case 01": {"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		"case 02": {"AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := FindRepeatedDNASeq(test.str)
			misc.Equals(t, test.ans, ans)
		})
	}
}
