package sol

/*
	LeetCode Problem 34: Find First and Last Position of Element in Sorted Array
	Level: Medium
	Description: Given an array of integers nums sorted in ascending order,
	find the starting and ending position of a given target value.
*/

// SearchRange find the starting and ending position of a given target value.
func SearchRange(nums []int, target int) []int {
	lower := LowerBound(nums, target)
	upper := LowerBound(nums, target+1)
	if lower < len(nums) && nums[lower] == target {
		return []int{lower, upper - 1}
	}
	return []int{-1, -1}
}

// LowerBound find the lowest index of a given target
func LowerBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r + l) / 2
		switch {
		case target > nums[m]:
			l = m + 1
		case target <= nums[m]:
			r = m - 1
		}
	}
	return l
}
