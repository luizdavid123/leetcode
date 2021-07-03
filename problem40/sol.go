package problem40

import "sort"

/*
	LeetCode Problem 40: Combination Sum II
	Level: Medium
	Description: Given a collection of candidate numbers (candidates) and a target number (target),
	find all unique combinations in candidates where the candidate numbers sum to target.
*/

// CombSumUnique return a list of all unique combinations of candidates where the chosen numbers sum to target.
func CombSumUnique(nums []int, tar int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	ch := make(chan []int, 1)
	go func(ch chan []int) {
		defer close(ch)
		BackTrack(ch, nums, []int{}, tar, 0)
	}(ch)
	for comb := range ch {
		ans = append(ans, comb)
	}
	return ans
}

// BackTrack is the helper method for finding all permutation
func BackTrack(ch chan []int, nums []int, comb []int, tar int, start int) {
	if tar == 0 {
		ch <- DeepCopyInts(comb)
		return
	}
	if tar < 0 {
		return
	}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		comb = append(comb, nums[i])
		BackTrack(ch, nums, comb, tar-nums[i], i+1)
		comb = comb[:len(comb)-1]
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
