package problem91

import (
	"leetcode/misc"
	"testing"
)

func TestNumOfDecodeWay(t *testing.T) {
	tests := map[string]struct {
		str string
		num int
	}{
		"case 01": {"12", 2},
		"case 02": {"226", 3},
		"case 03": {"0", 0},
		"case 04": {"06	", 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			num := NumOfDecodeWay(test.str)
			misc.Equals(t, test.num, num)
		})
	}
}

func TestValidCode(t *testing.T) {
	tests := map[string]struct {
		group string
		ok    bool
	}{
		"case 01": {"1", true},
		"case 02": {"12", true},
		"case 03": {"26", true},
		"case 04": {"0", false},
		"case 05": {"06", false},
		"case 06": {"27", false},
		"case 07": {"35", false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ok := ValidCode(test.group)
			misc.Equals(t, test.ok, ok)
		})
	}
}
