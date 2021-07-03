package problem45

import "math"

/*
	LeetCode Problem 45: Jump Game II
	Level: Medium
	Description: Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
	Each element in the array represents your maximum jump length at that position.
	Your goal is to reach the last index in the minimum number of jumps.
	You can assume that you can always reach the last index.
*/

// MinJumpV1 find the min step to reach the last index with recursion
func MinJumpV1(nums []int) int {
	return MinJumpByRecur(nums, 0)
}

// MinJumpByRecur find the min step to reach the last index at given step
func MinJumpByRecur(nums []int, pos int) int {
	if pos >= len(nums)-1 {
		return 0
	}
	min := math.MaxInt64 - 1
	for i := 1; i <= nums[pos]; i++ {
		min = MinInt(min, MinJumpByRecur(nums, pos+i)+1)
	}
	return min
}

// MinJumpV2 find the min step to reach the last index with recursion and memorizaton
func MinJumpV2(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt64 - 1
	}
	return MinJumpByRecurAndMem(nums, dp, 0)
}

// MinJumpByRecurAndMem find the min step to reach the last index at given step
func MinJumpByRecurAndMem(nums []int, dp []int, pos int) int {
	if pos >= len(nums)-1 {
		return 0
	}
	if dp[pos] != math.MaxInt64-1 {
		return dp[pos]
	}
	for i := 1; i <= nums[pos]; i++ {
		dp[pos] = MinInt(dp[pos], MinJumpByRecurAndMem(nums, dp, pos+i)+1)
	}
	return dp[pos]
}

// MinInt find the smaller int
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
