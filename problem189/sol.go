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

// RotateV2 rotate the array to the right by k steps
func RotateV2(nums []int, k int) []int {
	k %= len(nums)
	Reverse(nums, 0, len(nums)-1)
	Reverse(nums, 0, k-1)
	Reverse(nums, k, len(nums)-1)
	return nums
}

// Reverse reverse nums from start to end
func Reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
