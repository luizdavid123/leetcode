package problem41

/*
	LeetCode Problem 41: First Missing Positive
	Level: Medium
	Description: Given an unsorted integer array nums, find the smallest missing positive integer.
	You must implement an algorithm that runs in O(n) time and uses constant extra space.
*/

// FirstMissingPositive find the smallest missing positive integer
func FirstMissingPositive(nums []int) int {
	cnt := len(nums)
	for i := 0; i < cnt; i++ {
		for nums[i] > 0 && nums[i] <= cnt &&
			nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < cnt; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return cnt + 1
}
