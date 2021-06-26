package problem26

/*
	LeetCode Problem 26: Remove Duplicates from Sorted Array
	Level: Easy
	Description: Given an integer array nums sorted in non-decreasing order,
	remove the duplicates in-place such that each unique element appears only once.
	The relative order of the elements should be kept the same.
*/

// RmDup remove the duplicates in-place such that each unique element appears only once.
func RmDup(nums []int) ([]int, int) {
	if len(nums) == 0 {
		return []int{}, 0
	}
	u := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[u] && nums[i] > nums[u] && i != u {
			nums[i], nums[u+1] = nums[u+1], nums[i]
			u++
		}
	}
	return nums, u + 1
}
