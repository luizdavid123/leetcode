package problem53

import "math"

/*
	LeetCode Problem 53: Maximum Subarray
	Level: Easy
	Description: Given an integer array nums, find the contiguous subarray (containing at least one number)
	which has the largest sum and return its sum.
*/

// MaxSubSumV1 find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
func MaxSubSumV1(nums []int) int {
	sum, max := 0, math.MinInt64
	for _, num := range nums {
		sum += num
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
