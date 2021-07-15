package problem123

import "math"

/*
	LeetCode Problem 123: Best Time to Buy and Sell Stock III
	Level: Hard
	Description: You are given an array prices where prices[i] is the price of a given stock on the ith day.
	Find the maximum profit you can achieve. You may complete at most two transactions.
	Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
*/

// MaxProfit find the maximum profit when the number of txns can only be at most two
func MaxProfit(prices []int) int {
	buy_txn_01, buy_txn_02 := math.MinInt64, math.MinInt64
	sell_txn_01, sell_txn_02 := 0, 0
	for _, price := range prices {
		buy_txn_01 = MaxInt(buy_txn_01, -price)
		sell_txn_01 = MaxInt(sell_txn_01, buy_txn_01+price)
		buy_txn_02 = MaxInt(buy_txn_02, sell_txn_01-price)
		sell_txn_02 = MaxInt(sell_txn_02, buy_txn_02+price)

	}
	return MaxInt(0, sell_txn_02)
}

// MaxInt return the larger one
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
