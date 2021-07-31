package sol

import (
	"fmt"
	"sort"
	"strings"
)

/*
	LeetCode Problem 39: Combination Sum
	Level: Medium
	Description: Given an array of distinct integers candidates and a target integer target,
	return a list of all unique combinations of candidates where the chosen numbers sum to target.
	You may return the combinations in any order.
	The same number may be chosen from candidates an unlimited number of times. Two combinations are
	unique if the frequency of at least one of the chosen numbers is different.
	It is guaranteed that the number of unique combinations that sum up to target is less than 150
	combinations for the given input.
*/

// CombSum return a list of all unique combinations of candidates where the chosen numbers sum to target.
func CombSum(nums []int, tar int) [][]int {
	ans := [][]int{}
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
		comb = append(comb, nums[i])
		BackTrack(ch, nums, comb, tar-nums[i], i)
		comb = comb[:len(comb)-1]
	}
}

// BuildKeyFromComb build key from comb
func BuildKeyFromComb(comb []int) string {
	nums := comb
	sort.Ints(nums)
	counter := map[int]int{}
	for i := 0; i < len(nums); i++ {
		cnt, ok := counter[nums[i]]
		if !ok {
			counter[nums[i]] = 1
		} else {
			counter[nums[i]] = cnt + 1
		}
	}
	var sb strings.Builder
	for num, cnt := range counter {
		sb.WriteString(fmt.Sprintf("%d#%d", num, cnt))
	}
	return sb.String()
}

// DeepCopyInts return a copy of int slice
func DeepCopyInts(nums []int) []int {
	copy := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		copy[i] = nums[i]
	}
	return copy
}
