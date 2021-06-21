package problem7

import "math"

/*
	LeetCode Problem 7: Reverse Integer
	Level: Easy
	Description: Given a signed 32-bit integer x, return x with its digits reversed.
	If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1],
	then return 0.
*/

// ReverseNumber return n with its digits reversed.
func ReverseNumber(n int) int {
	ans := 0
	for d := n; d != 0; d = d / 10 {
		ans = ans*10 + d%10
		if ans < math.MinInt32 || ans > math.MaxInt32 {
			return 0
		}
	}
	return ans
}
