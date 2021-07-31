package sol

import (
	"leetcode/misc"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	tests := map[string]struct {
		colnum int
		title  string
	}{
		"case 01": {1, "A"},
		"case 02": {28, "AB"},
		"case 03": {701, "ZY"},
		"case 04": {2147483647, "FXSHRXW"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			title := ConvertToTitle(test.colnum)
			misc.Equals(t, test.title, title)
		})
	}
}

func TestGetTitleChar(t *testing.T) {
	tests := map[string]struct {
		num  int
		char string
	}{
		"case 01": {1, "A"},
		"case 02": {26, "Z"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			char := GetTitleChar(test.num)
			misc.Equals(t, test.char, char)
		})
	}
}
