package sol

/*
	LeetCode Problem 78: Subsets
	Level: Medium
	Description: Given an integer array nums of unique elements, return all possible subsets (the power set).
	The solution set must not contain duplicate subsets. Return the solution in any order.
*/

// GenSubSet return all possible subsets from an array of unique elements
func GenSubSet(nums []int) [][]int {
	ans := [][]int{}

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

// BackTrack is the helper method for GenSubset
func BackTrack(ch chan []int, nums []int, cand []int, start int) {
	ch <- DeepCopyInts(cand)
	for i := start; i < len(nums); i++ {
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
