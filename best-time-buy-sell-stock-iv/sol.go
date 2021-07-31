package sol

import "math"

/*
	LeetCode Problem 188: Best Time to Buy and Sell Stock IV
	Level: Hard
	Description: You are given an integer array prices where prices[i] is the price of a given stock on the ith day, and an integer k.
	Find the maximum profit you can achieve. You may complete at most k transactions.
	You may not engage in multiple transactions simultaneously
	Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
*/

// MaxProfit find the maximum profit you can achieve. You may complete at most k transactions.
func MaxProfit(prices []int, k int) int {
	buys := make([]int, k)
	sells := make([]int, k)
	for i := 0; i < k; i++ {
		buys[i] = math.MinInt64
		sells[i] = 0
	}
	for i := 0; i < len(prices); i++ {
		for j := 0; j < k; j++ {
			if j == 0 {
				buys[j] = MaxInt(buys[j], -prices[i])
				sells[j] = MaxInt(sells[j], buys[j]+prices[i])
			} else {
				buys[j] = MaxInt(buys[j], sells[j-1]-prices[i])
				sells[j] = MaxInt(sells[j], buys[j]+prices[i])
			}
		}
	}
	return MaxInt(0, sells[k-1])
}

// MaxInt return the larger one
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
