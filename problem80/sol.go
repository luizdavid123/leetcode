package problem80

/*
	LeetCode Problem 80: Remove Duplicates from Sorted Array II
	Level: Medium
	Description: Given an integer array nums sorted in non-decreasing order,
	remove some duplicates in-place such that each unique element appears at most twice.
	The relative order of the elements should be kept the same.
	You must do this by modifying the input array in-place with O(1) extra memory.
*/

// KeepUniqueNumAtMostTwo keep each unique number appears at most twice
func KeepUniqueNumAtMostTwo(nums []int) int {
	uni := 0
	limit := 2
	curr := nums[uni]
	for idx := 0; idx < len(nums); idx++ {
		if idx == 0 {
			limit--
			uni++
			continue
		}
		if nums[idx] == curr && limit > 0 {
			nums[uni] = nums[idx]
			limit--
			uni++
		} else if nums[idx] > curr {
			curr = nums[idx]
			nums[uni] = nums[idx]
			limit = 1
			uni++
		}
	}
	return uni
}
