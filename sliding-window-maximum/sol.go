package sol

import "math"

/*
	LeetCode Problem 239: Sliding Window Maximum
	Level: Hard
	Description: You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of
	the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves
	right by one position.
	Link: https://leetcode.com/problems/sliding-window-maximum/
*/

// MaxSlidingWindow return the max sliding window
func MaxSlidingWindowV2(nums []int, k int) []int {
	ans := []int{}
	for i := 0; i <= len(nums)-k; i++ {
		max := math.MinInt32
		ans = append(ans, max)
	}
	return ans
}

// MaxSlidingWindow return the max sliding window
func MaxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	for i := 0; i < len(nums)-k+1; i++ {
		win := nums[i : i+k]
		ans = append(ans, MaxIntInIns(win))
	}
	return ans
}

// maxIntInIns find max ele in eles
func MaxIntInIns(eles []int) int {
	cnt := len(eles)
	max := math.MinInt64
	for i := 0; i < cnt; i++ {
		if eles[i] > max {
			max = eles[i]
		}
	}
	return max
}

// MaxInt return the larger one
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
