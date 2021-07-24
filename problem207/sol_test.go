package problem207

import (
	"leetcode/misc"
	"testing"
)

func TestCanFinish(t *testing.T) {
	tests := map[string]struct {
		num  int
		pres [][]int
		ok   bool
	}{
		"case 01": {2, [][]int{{1, 0}}, true},
		"case 02": {2, [][]int{{1, 0}, {0, 1}}, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ok := CanFinish(test.num, test.pres)
			misc.Equals(t, test.ok, ok)
		})
	}
}
