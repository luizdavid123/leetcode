package problem213

/*
	LeetCode Problem 213: House Robber II
	Level: Medium
	Description: Given an integer array nums representing the amount of money of each house,
	return the maximum amount of money you can rob tonight without alerting the police.
	Link: https://leetcode.com/problems/house-robber-ii/
*/

// Rob return the maximum amount of money you can rob tonight without alerting the police
func Rob(nums []int) int {
	cnt := len(nums)
	if cnt == 0 {
		return 0
	}
	if cnt == 1 {
		return nums[0]
	}
	dpf, dps := make([]int, cnt+1), make([]int, cnt+1)
	dpf[0], dpf[1] = 0, nums[0]
	dps[0], dps[1] = 0, 0
	for i := 2; i <= cnt; i++ {
		dpf[i] = MaxInt(dpf[i-2]+nums[i-1], dpf[i-1])
		dps[i] = MaxInt(dps[i-2]+nums[i-1], dps[i-1])
	}
	return MaxInt(dpf[cnt-1], dps[cnt])
}

// MaxInt return the larger one
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
