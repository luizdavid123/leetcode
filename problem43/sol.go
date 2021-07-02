package problem43

import (
	"fmt"
	"strings"
)

/*
	LeetCode Problem 38: Multiply Strings
	Level: Medium
	Description: Given two non-negative integers num1 and num2 represented as strings,
	return the product of num1 and num2, also represented as a string.
*/

// Multiply return the product of num1 and num2, also represented as a string.
func Multiply(n string, m string) string {
	if n == "0" || m == "0" {
		return "0"
	}

	ncnt, mcnt := len(n), len(m)
	ans := make([]int, ncnt+mcnt)
	for i := range ans {
		ans[i] = 0
	}

	for i := ncnt - 1; i >= 0; i-- {
		for j := mcnt - 1; j >= 0; j-- {
			digit, curry := MultiplyDigit(n[i], m[j], ans[i+j+1])
			ans[i+j] += curry
			ans[i+j+1] = digit
		}
	}

	var sb strings.Builder
	for _, d := range ans {
		if sb.Len() != 0 || d != 0 {
			sb.WriteString(fmt.Sprint(d))
		}
	}
	return sb.String()
}

// MultiplyDigit Multiply digits one by one
func MultiplyDigit(l, r byte, c int) (int, int) {
	res := int((l-'0')*(r-'0')) + c
	digit, curry := res%10, res/10
	return digit, curry
}

// RvsStr reverse string
func RvsStr(str string) string {
	var rvs strings.Builder
	cnt := len(str)
	for i := cnt - 1; i >= 0; i-- {
		rvs.WriteByte(str[i])
	}
	return rvs.String()
}

// RvsIns reverse a list of int
func RvsIns(nums []int) []int {
	cnt := len(nums)
	for i := 0; i < cnt/2; i++ {
		nums[i], nums[cnt-1-i] = nums[cnt-1-i], nums[i]
	}
	return nums
}

// MinInt return the smaller int
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
