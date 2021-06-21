package problem8

import (
	"math"
	"strings"
)

/*
	LeetCode Problem 8: String to Integer (atoi)
	Level: Medium
	Description: Implement the myAtoi(string s) function,
	which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).
*/

// Atoi converts a string to a 32-bit signed integer
func Atoi(str string) int {
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return 0
	}

	first := str[0]
	pre := 1
	if first == '+' {
		pre = 1
		str = str[1:]
	}
	if first == '-' {
		pre = -1
		str = str[1:]
	}

	num := 0
	for _, ch := range str {
		if ch < '0' || ch > '9' {
			break
		}
		num = num*10 + (int(ch) - '0')
		if num > math.MaxInt32 {
			if pre == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
	}

	return pre * num
}
