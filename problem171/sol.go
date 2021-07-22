package problem171

import "math"

/*
	LeetCode Problem 171: Excel Sheet Column Number
	Level: Easy
	Description: Given a string columnTitle that represents the column title as appear in an Excel sheet,
	return its corresponding column number.
	Link: https://leetcode.com/problems/excel-sheet-column-number/
*/

// TitleNumber convert column title to column number
func TitleNumber(title string) int {
	cnt := len(title)
	ans := 0
	for i := 0; i < cnt; i++ {
		ch := title[i]
		pow := math.Pow(26, float64(cnt-i-1))
		ans += (int(ch) - 'A' + 1) * int(pow)
	}
	return ans
}
