package sol

import "math"

/*
	LeetCode Problem 33: Search in Rotated Sorted Array
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

		/*
		* Example: nums: [12, 13, 14, 15, 16, 17, 0, 1, 2, 3, 4, 5, 6, 7]
		* Left half: [12, 13, 14, 15, 16, 17] Right half: [0, 1, 2, 3, 4, 5, 6, 7]
		* If target is say in the left half, then when searching we need to make the numbers as
		* [12, 13, 14, 15, 16, 17, inf, inf, inf, inf, inf, inf, inf, inf]
		*	if target is in right half then we need to make it as
		* [-inf, -inf, -inf, -inf, -inf, -inf, 0, 1, 2, 3, 4, 5, 6, 7]
		 */

		// check if both targer and comp are on the same side
		if (target < nums[0]) == (comp < nums[0]) {
			// target and comp are on the same half
			comp = nums[m]
		} else {
			// target and comp are on the differet half
			// try to figure out where comp is and making comp as -inf or inf
			if target < nums[0] {
				// target is on the right half, comp is on the left half
				// make comp to -inf
				comp = math.MinInt32
			} else {
				// target is on the left half, comp is on the right half
				// make comp to +inf
				comp = math.MaxInt32
			}
		}

		// normal bineary search procedure
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
