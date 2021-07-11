package problem71

import (
	"leetcode/misc"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := map[string]struct {
		path string
		out  string
	}{
		"case 01": {"/home/", "/home"},
		"case 02": {"/../", "/"},
		"case 03": {"/home//foo/", "/home/foo"},
		"case 04": {"/a/./b/../../c/", "/c"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := SimplifyPath(test.path)
			misc.Equals(t, test.out, out)
		})
	}
}
