package sol

/*
	LeetCode Problem 125: Climbing Stairs
	Level: Easy
	Description: You are climbing a staircase. It takes n steps to reach the top.
	Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/

// ClimbStairsV1 return the num of way climt to the top by recursion
func ClimbStairsV1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return ClimbStairsV1(n-1) + ClimbStairsV1(n-2)
}

// ClimbStairsV2 return the num of way climt to the top memorization
func ClimbStairsV2(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = -1
	}
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
