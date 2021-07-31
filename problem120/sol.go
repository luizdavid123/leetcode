package sol

import (
	"math"
)

/*
	LeetCode Problem 120: Triangle
	Level: Medium
	Description: Given a triangle array, return the minimum path sum from top to bottom.
*/

// MinTriPathSum return the minimum path sum from top to bottom of tri arr
func MinTriPathSum(tri [][]int) int {
	nrow := len(tri)
	ncol := len(tri[nrow-1])
	dp := make([]int, ncol)
	dp[0] = tri[0][0]
	for i := 1; i < nrow; i++ {
		row := tri[i]
		ncell := len(row)
		for j := ncell - 1; j >= 0; j-- {
			switch j {
			case 0:
				dp[j] = dp[j] + row[j]
			case ncell - 1:
				dp[j] = dp[j-1] + row[j]
			default:
				dp[j] = MinInt(dp[j-1], dp[j]) + row[j]
			}
		}
	}
	return MinIntInInts(dp)
}

// MinInt return the smaller one
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// MinIntInInts find min ele in eles
func MinIntInInts(eles []int) int {
	cnt := len(eles)
	min := math.MaxInt64
	for i := 0; i < cnt; i++ {
		if eles[i] < min {
			min = eles[i]
		}
	}
	return min
}
