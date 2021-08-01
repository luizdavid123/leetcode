package sol

import "strings"

/*
	LeetCode Problem 214: Shortest Palindrome
	Level: Hard
	Description: You are given a string s. You can convert s to a palindrome by adding characters in front of it.
	Return the shortest palindrome you can find by performing this transformation.
	Link: https://leetcode.com/problems/shortest-palindrome/
*/

// ShortestPalindrome return the shortest palindrome you can find by performing this transformation.
func ShortestPalindrome(str string) string {
	j := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == str[j] {
			j++
		}
	}
	if j == len(str) {
		return str
	}
	suffix := str[j:]
	return RvsStr(suffix) + ShortestPalindrome(str[:j]) + suffix
}

// RvsStr reverse str
func RvsStr(str string) string {
	var rvs strings.Builder
	cnt := len(str)
	for i := cnt - 1; i >= 0; i-- {
		rvs.WriteByte(str[i])
	}
	return rvs.String()
}
