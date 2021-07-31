package sol

import (
	"leetcode/misc"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := map[string]struct {
		str string
		ans [][]string
	}{
		"case 01": {"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		"case 02": {"a", [][]string{{"a"}}},
		"case 03": {"", [][]string{{}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := Partition(test.str)
			misc.Equals(t, test.ans, ans)
		})
	}
}
