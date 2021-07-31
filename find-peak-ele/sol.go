package sol

/*
	LeetCode Problem 162: Find Peak Element
	Level: Medium
	Description: A peak element is an element that is strictly greater than its neighbors.
	Given an integer array nums, find a peak element, and return its index.
	If the array contains multiple peaks, return the index to any of the peaks
	Link: https://leetcode.com/problems/find-peak-element/
*/

// FindPeakEle find a peak element and return its index
func FindPeakEle(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
