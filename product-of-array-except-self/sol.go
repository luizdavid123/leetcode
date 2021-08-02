package sol

/*
	LeetCode Problem 238: Product of Array Except Self
	Level: Medium
	Description: Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the
	elements of nums except nums[i].
	Link: https://leetcode.com/problems/product-of-array-except-self/
*/

// ProductExceptSelf return an array answer such that answer[i] is equal to the product of all the
// elements of nums except nums[i].
func ProductExceptSelf(nums []int) []int {
	cnt := len(nums)

	left := make([]int, cnt)
	left[0] = 1
	for i := 1; i < cnt; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := make([]int, cnt)
	right[cnt-1] = 1
	for i := cnt - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	ans := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}
