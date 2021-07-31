package sol

/*
	LeetCode Problem 47: Permutations
	Level: Medium
	Description: Given an array nums of distinct integers, return all the possible permutations.
	You can return the answer in any order.
*/

/*
	As a reminder, backtracking is a general algorithm for finding all (or some) solutions to some problems
	with constraints. It incrementally builds candidates to the solutions, and abandons a candidate as soon
	as it determines that the candidate cannot possibly lead to a solution.
*/

// Perm return all the possible permutations
func Perm(nums []int) [][]int {
	ans := [][]int{}
	ch := make(chan []int)
	go func(ch chan []int) {
		defer close(ch)
		BackTrack(ch, nums, []int{})
	}(ch)
	for comb := range ch {
		ans = append(ans, comb)
	}
	return ans
}

// BackTrack is the helper method for finding all permutation
func BackTrack(ch chan []int, nums []int, comb []int) {
	ncomb, nnums := len(comb), len(nums)
	if ncomb == nnums {
		ch <- DeepCopyInts(comb)
		return
	}
	for i := 0; i < nnums; i++ {
		if IsIntInInts(comb, nums[i]) {
			continue
		}
		comb = append(comb, nums[i])
		BackTrack(ch, nums, comb)
		comb = comb[:len(comb)-1]
	}
}

// IsIntInInts check if int is in ins
func IsIntInInts(nums []int, num int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return true
		}
	}
	return false
}

// DeepCopyInts return the deepcopy of a slice of int
func DeepCopyInts(nums []int) []int {
	cnt := len(nums)
	copy := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		copy[i] = nums[i]
	}
	return copy
}
