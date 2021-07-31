package sol

/*
	LeetCode Problem 217: Contains Duplicate
	Level: Easy
	Description: Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
	Link: https://leetcode.com/problems/contains-duplicate/
*/

// HasDup check if any values in arr appears >= 2 time
func HasDup(nums []int) bool {
	seen := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if seen[nums[i]] >= 1 {
			return true
		}
		seen[nums[i]] += 1
	}
	return false
}
