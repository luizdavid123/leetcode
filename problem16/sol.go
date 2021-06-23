package problem16

import (
	"math"
	"sort"
)

/*
	LeetCode Problem 16:  3Sum Closest
	Level: Medium
	Description: Given an array nums of n integers and an integer target,
	find three integers in nums such that the sum is closest to target.
	Return the sum of the three integers. You may assume that each input would have exactly one solution.
*/

// ThreeSumClosest find three integers in nums such that the sum is closest to target.
func ThreeSumClosest(nums []int, target int) int {
	min, max, cnt := math.MaxInt16, 0, len(nums)

	sort.Ints(nums)

	for i := 0; i < cnt-2; i++ {
		l, r := i+1, cnt-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum == target:
				return sum
			case sum < target:
				l++
			case sum > target:
				r--
			}

			diff := AbsInt(target - sum)
			if diff < min {
				min = diff
				max = sum
			}
		}
	}

	return max
}

// AbsInt return absoulte value
func AbsInt(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

// MinInt return smaller integer
func MinInt(l, r int) int {
	if l < r {
		return l
	}
	return r
}
