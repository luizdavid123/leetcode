package problem122

/*
	LeetCode Problem 122: Best Time to Buy and Sell Stock II
	Level: Easy
	Description: You are given an array prices where prices[i] is the price of a given stock on the ith day.
	Find the maximum profit you can achieve. You may complete as many transactions as you like
	Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
*/

// MaxProfit find the maximum profit when the nums of txn is not limited
func MaxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
