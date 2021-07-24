package problem205

import "fmt"

/*
	LeetCode Problem 205: Isomorphic Strings
	Level: Easy
	Description: Given two strings s and t, determine if they are isomorphic.
	Link: https://leetcode.com/problems/isomorphic-strings/
*/

// IsIsomorphic determine if two strs are isomorphic.
func IsIsomorphic(s, t string) bool {
	return CharToIdx(s) == CharToIdx(t)
}

// CharToIdx replace char in str with its idx
func CharToIdx(str string) string {
	buf := ""
	dict := map[byte]int{}
	for i := 0; i < len(str); i++ {
		_, ok := dict[str[i]]
		if !ok {
			dict[str[i]] = i
		}
		buf += fmt.Sprint(dict[str[i]])
		buf += " "
	}
	return buf
}
