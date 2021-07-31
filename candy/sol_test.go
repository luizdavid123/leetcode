package sol

import (
	"leetcode/misc"
	"testing"
)

func TestMinCandy(t *testing.T) {
	tests := map[string]struct {
		ratings []int
		ans     int
	}{
		"case 01": {[]int{1, 0, 2}, 5},
		"case 02": {[]int{1, 2, 2}, 4},
		"case 03": {[]int{}, 0},
		"case 04": {[]int{1}, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ans := MinCandy(test.ratings)
			misc.Equals(t, test.ans, ans)
		})
	}
}
