package sol

import (
	"strings"
)

/*
	LeetCode Problem 125: Valid Palindrome
	Level: Easy
	Description: Given a string s, determine if it is a palindrome,
	considering only alphanumeric characters and ignoring cases.
*/

// IsPalindrome check if it is a palindrome
func IsPalindrome(str string) bool {
	cand := Normalize(str)
	cnt := len(cand)
	for i := 0; i < cnt/2; i++ {
		if cand[i] != cand[cnt-i-1] {
			return false
		}
	}
	return true
}

// Normalize return a string containing only letters of lower case or digits
func Normalize(str string) string {
	var builder strings.Builder
	for i := 0; i < len(str); i++ {
		if IsDigitOrLetter(str[i]) {
			builder.WriteString(strings.ToLower(str[i : i+1]))
		}
	}
	return builder.String()
}

// IsDigitOrLetter check if it is a letter or a digit
func IsDigitOrLetter(ch byte) bool {
	switch {
	case ch >= 'a' && ch <= 'z':
		return true
	case ch >= 'A' && ch <= 'Z':
		return true
	case ch >= '0' && ch <= '9':
		return true
	default:
		return false
	}
}
