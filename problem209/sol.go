package problem209

import "math"

/*
	LeetCode Problem 209:  Minimum Size Subarray Sum
	Level: Medium
	Description: Given an array of positive integers nums and a positive integer target, return the minimal length of a
	contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to
	target. If there is no such subarray, return 0 instead.
	Link: https://leetcode.com/problems/minimum-size-subarray-sum/
*/

// MinSubArrSumLen find the min len of contiguous subarr of which the sum >= tar
func MinSubArrSumLen(nums []int, tar int) int {
	ans := math.MaxInt64
	left, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= tar {
			ans = MinInt(ans, i-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// MinInt return the smaller one
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
