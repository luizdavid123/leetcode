package problem168

import "strings"

/*
	LeetCode Problem 168: Excel Sheet Column Title
	Level: Easy
	Description: Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.
	Link: https://leetcode.com/problems/excel-sheet-column-title/
*/

// ConvertToTitle return its corresponding column title as it appears in an Excel sheet given
// an integer colnum
func ConvertToTitle(col int) string {
	ans := ""
	d := col
	for d > 0 {
		d--
		ans += GetTitleChar(d % 26)
		d /= 26
	}
	return RvsStr(ans)
}

// GetTitleChar return single char between A-Z based on the given num in range 1-26
func GetTitleChar(num int) string {
	chars := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return chars[num]
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
