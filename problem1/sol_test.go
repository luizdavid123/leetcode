package problem1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct {
		nums []int
		tar  int
		want []int
	}{
		"case 01": {nums: []int{2, 7, 11, 15}, tar: 9, want: []int{0, 1}},
		"case 02": {nums: []int{3, 2, 4}, tar: 6, want: []int{1, 2}},
		"case 03": {nums: []int{3, 3}, tar: 6, want: []int{0, 1}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := HashMap(test.nums, test.tar)
			assert.Equal(t, test.want, got)
		})
	}
}
