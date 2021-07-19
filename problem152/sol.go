package problem152

/*
	LeetCode Problem 152: Maximum Product Subarray
	Level: Medium
	Description: Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.
	Link: https://leetcode.com/problems/maximum-product-subarray/
*/

// MaxProduct find a contiguous non-empty subarray within the array that has the largest product
func MaxProduct(nums []int) int {
	ans, min, max := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}
		max = MaxInt(nums[i], max*nums[i])
		min = MinInt(nums[i], min*nums[i])
		ans = MaxInt(ans, max)
	}
	return ans
}

// MaxInt return the larger one
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt return the smaller one
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
