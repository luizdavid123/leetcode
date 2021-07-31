package sol

import "math"

/*
	LeetCode Problem 29: Divide Two Integers
	Level: Easy
	Description: Given two integers dividend and divisor,
	divide two integers without using multiplication, division, and mod operator.
*/

// Divide divide two integers without using multiplication, division, and mod operator
func Divide(dividend int, divisor int) int {
	if dividend <= math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	pos := 1
	if dividend < 0 {
		dividend *= -1
		pos *= -1
	}
	if divisor < 0 {
		divisor *= -1
		pos *= -1
	}

	quo := 0
	for dividend >= divisor && dividend > 0 {
		dividend -= divisor
		quo++
	}

	return quo * pos
}
