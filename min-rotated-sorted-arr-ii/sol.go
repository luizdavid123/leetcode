package sol

/*
	LeetCode Problem 153: Find Minimum in Rotated Sorted Array II
	Level: Medium
	Description: Given the sorted rotated array nums of unique elements, return the minimum element of this array.
	You must write an algorithm that runs in O(log n) time.
	Link: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
*/

// FindMinInRotatedArr find the min num in the rotated array with duplicates
func FindMinInRotatedArr(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[r] {
			l = m + 1
		} else if nums[m] < nums[r] {
			r = m
		} else {
			r--
		}
	}
	return nums[l]
}

// FindMinInRotatedArrV2 find the min num in the rotated array with duplicates
func FindMinInRotatedArrV2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for nums[l] == nums[r] && l < r {
		r--
	}
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[m+1] {
			return nums[m+1]
		}
		if nums[m-1] > nums[m] {
			return nums[m]
		}
		if nums[m] > nums[0] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
