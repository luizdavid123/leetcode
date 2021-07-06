package problem90

import "sort"

/*
	LeetCode Problem 78: Subsets II
	Level: Medium
	Description: Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
	The solution set must not contain duplicate subsets. Return the solution in any order.
*/

// GenSubSetWithDup return all possible subsets from an array of unique elements
func GenSubSetWithDup(nums []int) [][]int {
	ans := [][]int{}

	sort.Ints(nums)

	ch := make(chan []int)
	go func(ch chan []int) {
		defer close(ch)
		BackTrack(ch, nums, []int{}, 0)
	}(ch)

	for cand := range ch {
		ans = append(ans, cand)
	}

	return ans
}

// BackTrack is the helper method for GenSubsetWithDup
func BackTrack(ch chan []int, nums []int, cand []int, start int) {
	ch <- DeepCopyInts(cand)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		cand = append(cand, nums[i])
		BackTrack(ch, nums, cand, i+1)
		cand = cand[:len(cand)-1]
	}
}

// DeepCopyInts return a copy of int slice
func DeepCopyInts(nums []int) []int {
	copy := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		copy[i] = nums[i]
	}
	return copy
}
