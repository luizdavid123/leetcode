package problem17

import "sort"

/*
	LeetCode Problem 18:  4Sum
	Level: Medium
	Description: Given an array nums of n integers, return an array of all the unique quadruplets
	[nums[a], nums[b], nums[c], nums[d]] such that:
	1. 0 <= a, b, c, d < n
	2. a, b, c, and d are distinct
	3. nums[a] + nums[b] + nums[c] + nums[d] == target
*/

// FourSum finds all distinct quadruplets in nums that sum up to tager
func FourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	cnt := len(nums)

	sort.Ints(nums)

	for i := 0; i < cnt-3; i++ {
		for j := i + 1; j < cnt-2; j++ {
			l, r := j+1, cnt-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case sum == target:
					quad := []int{nums[i], nums[j], nums[l], nums[r]}
					sort.Ints(quad)
					if !IsAnsIncludesQuad(ans, quad) {
						ans = append(ans, quad)
					}
					l++
					r--
				case sum < target:
					l++
				case sum > target:
					r--
				}
			}
		}
	}

	return ans
}

// IsAnsIncludesQuad check if ans contains ans
func IsAnsIncludesQuad(ans [][]int, quad []int) bool {
	for i := 0; i < len(ans); i++ {
		arr := ans[i]
		if arr[0] == quad[0] &&
			arr[1] == quad[1] &&
			arr[2] == quad[2] &&
			arr[3] == quad[3] {
			return true
		}
	}
	return false
}
