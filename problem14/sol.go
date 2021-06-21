package problem14

import "strings"

/*
	LeetCode Problem 14: Longest Common Prefix
	Level: Easy
	Description: Write a function to find the longest common prefix string amongst an array of strings.
*/

// LongestCommonPrefix find the longest common prefix string amongst an array of strings
func LongestCommonPrefix(strs []string) string {
	max := MaxLenStrInStrs(strs)
	for len(max) > 0 {
		passed := 0
		for i := 0; i < len(strs); i++ {
			if strings.Index(strs[i], max) != 0 {
				max = max[:len(max)-1]
				passed = 0
				break
			} else {
				passed++
			}
		}
		if passed == len(strs) {
			break
		}
	}
	return max
}

// MaxLenStrInStrs return a string with maximum length amongest an array of strings
func MaxLenStrInStrs(strs []string) string {
	max := ""
	for i := 0; i < len(strs); i++ {
		if len(max) < len(strs[i]) {
			max = strs[i]
		}
	}
	return max
}
