package paranumber

import "math"

/*
	LeetCode Problem 9 Palindrome Number
	Level: Easy
	Description: Given an integer x, return true if x is palindrome integer.
*/

// IsPaliNum check if n is palindrome interger
func IsPaliNum(n int) bool {
	if n < 0 {
		return false
	}
	r := ReverseNumber(n)
	return r == n
}

// ReverseNumber return n with its digits reversed.
func ReverseNumber(n int) int {
	r := 0
	for d := n; d != 0; d = d / 10 {
		r = r*10 + d%10
		if r < math.MinInt32 || r > math.MaxInt32 {
			return 0
		}
	}
	return r
}
