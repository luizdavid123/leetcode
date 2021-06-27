package problem30

import (
	"leetcode/misc"
	"testing"
)

func TestFindSubStr(t *testing.T) {
	tests := map[string]struct {
		str   string
		words []string
		idxs  []int
	}{
		"case 01": {"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		"case 02": {"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		"case 03": {"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			idxs := FindSubStr(test.str, test.words)
			misc.Equals(t, test.idxs, idxs)
		})
	}
}

func TestIsTwoStrIntMapEqual(t *testing.T) {
	tests := map[string]struct {
		left  map[string]int
		right map[string]int
		equal bool
	}{
		"case 01": {map[string]int{}, map[string]int{}, true},
		"case 02": {map[string]int{"foo": 1}, map[string]int{"foo": 1}, true},
		"case 03": {map[string]int{"foo": 1, "bar": 2}, map[string]int{"bar": 2, "foo": 1}, true},
		"case 04": {map[string]int{"foo": 1}, map[string]int{}, false},
		"case 05": {map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1}, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			equal := IsTwoStrIntMapEqual(test.left, test.right)
			misc.Equals(t, test.equal, equal)
		})
	}
}
