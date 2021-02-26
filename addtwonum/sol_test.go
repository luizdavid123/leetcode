package addtwonum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func SeedList() *ListNode {
	tmp := NewListNode(0)
	ptrt := tmp
	for i := 1; i < 10; i++ {
		ptrt.Next = NewListNode(i)
		ptrt = ptrt.Next
	}
	return tmp
}

func TestAddTwoNumbers(t *testing.T) {
	tests := map[string]struct {
		left  string
		right string
		exp   string
	}{
		"same length":                  {left: "0001", right: "0002", exp: "0003"},
		"same length with one carry":   {left: "1001", right: "9002", exp: "0103"},
		"same length with all carry":   {left: "1001", right: "9992", exp: "0004"},
		"same length with last carry":  {left: "1001", right: "9998", exp: "00001"},
		"right longer":                 {left: "001", right: "0002", exp: "0012"},
		"right longer with one carry":  {left: "011", right: "0902", exp: "0022"},
		"right longer with all carry":  {left: "111", right: "9992", exp: "0113"},
		"right longer with last carry": {left: "111", right: "9999", exp: "01101"},
		"left longer":                  {right: "001", left: "0002", exp: "0012"},
		"left longer with one carry":   {right: "011", left: "0902", exp: "0022"},
		"left longer with all carry":   {right: "111", left: "9992", exp: "0113"},
		"left longer with last carry":  {right: "111", left: "9999", exp: "01101"},
		"left with length one":         {right: "1", left: "9999", exp: "00001"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			left := NewListFromString(test.left)
			right := NewListFromString(test.right)
			got := AddTwoNumbers(left, right)
			assert.Equal(t, test.exp, got.String())
		})
	}
}

func Test_NewListFromString(t *testing.T) {
	str := "0123456789"
	tmp := NewListFromString(str)
	for i := 0; i < 9; i++ {
		assert.Equal(t, i, tmp.Val)
		tmp = tmp.Next
	}
}

func Test_String(t *testing.T) {
	tmp := SeedList()
	exp := "0123456789"
	assert.Equal(t, exp, tmp.String())
}

func Test_DeepCopy(t *testing.T) {
	tmp := SeedList()
	copy := tmp.DeepCopy()

	ptrc := copy
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, ptrc.Val)
		ptrc = ptrc.Next
	}

	for copy.Next != nil && tmp != nil {
		assert.NotSame(t, tmp, copy)
		copy = copy.Next
		tmp = tmp.Next
	}
}
