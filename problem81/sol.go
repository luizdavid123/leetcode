package sol

import "math"

/*
	LeetCode Problem 81: Search in Rotated Sorted Array II
	Level: Medium
	Description: Given the array nums after the rotation and an integer target,
	return true if target is in nums, or false if it is not in nums.
*/

// SearchInRotatedArr find the target in the rotated array
func SearchInRotatedArr(nums []int, tar int) bool {
	l, r := 0, len(nums)-1
	for l <= r {

		// To avoid duplicate
		for l < r && nums[l] == nums[l+1] {
			l++
		}
		for l < r && nums[r] == nums[r-1] {
			r--
		}

		m := (r-l)/2 + l
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
		if (tar < nums[0]) == (comp < nums[0]) {
			// target and comp are on the same half
			comp = nums[m]
		} else {
			// target and comp are on the differet half
			// try to figure out where comp is and making comp as -inf or inf
			if tar < nums[0] {
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
		case tar == comp:
			return true
		case tar > comp:
			l = m + 1
		case tar < comp:
			r = m - 1
		}
	}
	return false
}
