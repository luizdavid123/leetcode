package problem121

import "math"

/*
	LeetCode Problem 121: Best Time to Buy and Sell Stock
	Level: Easy
	Description: You are given an array prices where prices[i] is the price of a given stock on the ith day.
	You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
	Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
	Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/

// MaxProfitBruteForce return the maximum profit you can achieve from this transaction
func MaxProfitBruteForce(prices []int) int {
	ans := 0
	for buy := 0; buy < len(prices); buy++ {
		for sell := buy + 1; sell < len(prices); sell++ {
			profit := prices[sell] - prices[buy]
			ans = MaxInt(ans, profit)
		}
	}
	return ans
}

// MaxProfitV2 return the maximum profit you can achieve from this transaction
func MaxProfitV2(prices []int) int {
	ans := 0
	min := math.MaxInt64
	for _, price := range prices {
		if min > price {
			min = price
		} else if ans < price-min {
			ans = price - min
		}
	}
	return ans
}

// MaxProfitV3 return the maximum profit you can achieve from this transaction
func MaxProfitV3(prices []int) int {
	sell, buy := 0, math.MinInt64
	for _, price := range prices {
		buy = MaxInt(buy, -price)
		sell = MaxInt(sell, buy+price)
	}
	return sell
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
