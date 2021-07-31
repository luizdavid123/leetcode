package sol

import "strings"

/*
	LeetCode Problem 71: Simplify Path
	Level: Medium
	Description: Given a string path, which is an absolute path (starting with a slash '/') to a
	file or directory in a Unix-style file system, convert it to the simplified canonical path.
*/

// SimplifyPath convert an absolute path to a canonical path
func SimplifyPath(path string) string {
	stack := []string{}
	ignore := []string{"..", ".", ""}
	for _, dir := range strings.Split(path, "/") {
		if dir == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if !IsStrInStrs(ignore, dir) {
			stack = append(stack, dir)
		}
	}
	res := ""
	for i := len(stack) - 1; i >= 0; i-- {
		dir := stack[i]
		res = "/" + dir + res
	}
	if res == "" {
		return "/"
	}
	return res
}

// IsStrInStrs check if an element is in an array
func IsStrInStrs(eles []string, ele string) bool {
	cnt := len(eles)
	for i := 0; i < cnt; i++ {
		if eles[i] == ele {
			return true
		}
	}
	return false
}
