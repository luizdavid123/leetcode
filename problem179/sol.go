package problem179

import (
	"sort"
	"strconv"
)

/*
	LeetCode Problem 179: Largest Number
	Level: Medium
	Description: Given a list of non-negative integers nums, arrange them such that they form the largest number.
	Link: https://leetcode.com/problems/largest-number/
*/

type SortBy []string

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i]+a[j] > a[j]+a[i] }

/*
	Insight:
	lhs, rhs = 31, 9
	lhs+rhs, rhs+lhs = 319, 931
	since 931 > 319, rhs should put in front of lhs
*/

// LargestNumber return the largest number by arranging nums
func LargestNumber(nums []int) string {
	cnt := len(nums)
	strs := make([]string, cnt)
	for i := 0; i < cnt; i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(SortBy(strs))
	if strs[0][0] == '0' {
		return "0"
	}
	ans := ""
	for i := 0; i < cnt; i++ {
		ans += strs[i]
	}
	return ans
}
