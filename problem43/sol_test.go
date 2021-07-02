package problem43

import (
	"leetcode/misc"
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := map[string]struct {
		n string
		m string
		o string
	}{
		"case 01": {"1", "5", "5"},
		"case 02": {"2", "5", "10"},
		"case 03": {"3", "15", "45"},
		"case 04": {"15", "5", "75"},
		"case 05": {"123", "456", "56088"},
		"case 06": {"0", "123", "0"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			o := Multiply(test.n, test.m)
			misc.Equals(t, test.o, o)
		})
	}
}

func TestRvsStr(t *testing.T) {
	tests := map[string]struct {
		o string
		r string
	}{
		"case 01": {"1", "1"},
		"case 02": {"12", "21"},
		"case 03": {"123", "321"},
		"case 04": {"", ""},
		"case 05": {"1234", "4321"},
		"case 06": {"12345", "54321"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := RvsStr(test.o)
			misc.Equals(t, test.r, r)
		})
	}
}
