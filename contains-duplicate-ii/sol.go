package sol

/*
	LeetCode Problem 219: Contains Duplicate II
	Level: Easy
	Description: Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the
	array such that nums[i] == nums[j] and abs(i - j) <= k.
	Link: https://leetcode.com/problems/contains-duplicate-ii/
*/

// ContainsNearbyDuplcate check if there there are two distinct indices i and j in the
// array such that nums[i] == nums[j] and abs(i - j) <= k.
func ContainsNearbyDuplcate(nums []int, k int) bool {
	seen := map[int]int{}
	for i := 0; i < len(nums); i++ {
		idx, ok := seen[nums[i]]
		if !ok {
			seen[nums[i]] = i
		} else {
			if i-idx <= k {
				return true
			} else {
				seen[nums[i]] = i
			}
		}
	}
	return false
}

// MinInt return the smaller one
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
