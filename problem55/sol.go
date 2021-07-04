package problem55

/*
	LeetCode Problem 55: Jump Game
	Level: Easy
	Description: Given an array of non-negative integers nums,
	you are initially positioned at the first index of the array.
	Each element in the array represents your maximum jump length at that position.
	Determine if you are able to reach the last index.
*/

// Helper Data Structure for memorizaition
type DP struct {
	Can     bool
	Checked bool
}

// CanJumpV1 determine if it is able to reach the last index at the first index of the array
// CanJumpV1 solve it by using recursion and memorization
func CanJumpV1(nums []int) bool {
	dp := make([]DP, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i].Can = false
		dp[i].Checked = false
	}
	return CanJumpRecurNMem(nums, dp, 0)
}

//  CanJumpRecurNMem is the helper method to check if it is able to reach the last index at
// 	index pos
func CanJumpRecurNMem(nums []int, dp []DP, pos int) bool {
	if pos >= len(nums)-1 {
		return true
	}
	if dp[pos].Checked {
		return dp[pos].Can
	}
	for i := 1; i <= nums[pos]; i++ {
		dp[pos].Can = dp[pos].Can || CanJumpRecurNMem(nums, dp, i+pos)
		dp[pos].Checked = true
	}
	return dp[pos].Can
}

// CanJumpV2 determine if it is able to reach the last index at the first index of the array
// CanJumpV2 solve it by Greedy approach
func CanJumpV2(nums []int) bool {
	far := nums[0]
	for i := 1; i < len(nums); i++ {
		if far < i {
			return false
		}
		far = MaxInt(far, nums[i]+i)
		if far >= len(nums)-1 {
			return true
		}
	}
	return true
}

// MaxInt return the larger int
func MaxInt(n, m int) int {
	if n > m {
		return n
	}
	return m
}
