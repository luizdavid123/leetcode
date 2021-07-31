package sol

import (
	"leetcode/misc"
	"testing"
)

func TestTitleNumber(t *testing.T) {
	tests := map[string]struct {
		title string
		num   int
	}{
		"case 01": {"A", 1},
		"case 02": {"AB", 28},
		"case 03": {"ZY", 701},
		"case 04": {"FXSHRXW", 2147483647},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			num := TitleNumber(test.title)
			misc.Equals(t, test.num, num)
		})
	}
}
