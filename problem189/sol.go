package problem189

/*
	LeetCode Problem 189: Rotate Array
	Level: Medium
	Description: Given an array, rotate the array to the right by k steps, where k is non-negative.
	Link: https://leetcode.com/problems/rotate-array/
*/

// Rotate rotate the array to the right by k steps
func Rotate(nums []int, k int) []int {
	cnt := len(nums)
	ans := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		ans[(i+k)%cnt] = nums[i]
	}
	return ans
}
