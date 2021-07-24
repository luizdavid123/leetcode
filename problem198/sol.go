package problem198

/*
	LeetCode Problem 198: House Robber
	Level: Medium
	Description: You are a professional robber planning to rob houses along a street. Each house has a certain amount of money
	stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems
	connected and it will automatically contact the police if two adjacent houses were broken into on the same
	night.
	Given an integer array nums representing the amount of money of each house, return the maximum amount of
	money you can rob tonight without alerting the police.
	Link: https://leetcode.com/problems/house-robber/
*/

// MaxRob find the maximum amount of money you can rob tonight without alerting the police.
func MaxRob(nums []int) int {
	cnt := len(nums)
	dp := make([]int, cnt+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 1; i < cnt; i++ {
		dp[i+1] = MaxInt(dp[i], dp[i-1]+nums[i])
	}
	return dp[cnt]
}

// MaxInt return the larger one
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
