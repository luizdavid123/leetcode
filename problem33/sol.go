package problem33

import "math"

/*
	LeetCode Problem 30: Search in Rotated Sorted Array
	Level: Medium
	Description: There is an integer array nums sorted in ascending order (with distinct values).
	Prior to being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
	such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
	Given the array nums after the rotation and an integer target,
	return the index of target if it is in nums, or -1 if it is not in nums.
*/

// FindIntInRotatedInts find the targer in the sorted and rotated array
// Ref: https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/154836/The-INF-and-INF-method-but-with-a-better-explanation-for-dummies-like-me
func FindIntInRotatedInts(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (r + l) / 2
		comp := nums[m]

		// check if both targer and comp are on the same side
		if (target < nums[0]) == (comp < nums[0]) {
			comp = nums[m]
		} else {
			// try to figure out where comp is and making comp as -inf or inf
			if target < nums[0] {
				comp = math.MinInt32
			} else {
				comp = math.MaxInt32
			}
		}

		switch {
		case target == comp:
			return m
		case target > comp:
			l = m + 1
		case target < comp:
			r = m
		}
	}
	return -1
}
