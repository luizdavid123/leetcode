package problem15

import (
	"sort"
)

/*
	LeetCode Problem 15: 3Sum
	Level: Medium
	Description: Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
	such that i != j, i != k,and j != k, and nums[i] + nums[j] + nums[k] == 0.
*/

// ThreeSum return all return all the triplets [nums[i], nums[j], nums[k]] that sum up to 0
func ThreeSum(nums []int, tar int) [][]int {
	ans := [][]int{}

	cnt := len(nums)
	sorted := nums
	sort.Ints(nums)

	for i := 0; i < cnt-2; i++ {
		l, r := i+1, cnt-1
		for l < r {
			sum := sorted[i] + sorted[l] + sorted[r]
			switch {
			case sum == tar:
				triplet := []int{sorted[i], sorted[l], sorted[r]}
				sort.Ints(triplet)
				if !IsAnsIncludesTripe(ans, triplet) {
					ans = append(ans, triplet)
				}
				l++
				r--
			case sum < tar:
				l++
			case sum > tar:
				r--
			}
		}
	}

	return ans
}

// IsAnsIncludesTripe check if ans contains ans
func IsAnsIncludesTripe(ans [][]int, triplet []int) bool {
	for i := 0; i < len(ans); i++ {
		arr := ans[i]
		if arr[0] == triplet[0] &&
			arr[1] == triplet[1] &&
			arr[2] == triplet[2] {
			return true
		}
	}
	return false
}
