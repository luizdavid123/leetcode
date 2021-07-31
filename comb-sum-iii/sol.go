package sol

/*
	LeetCode Problem 216: Combinaton Sum III
	Level: Mediumn
	Description: Find all valid combinations of k numbers that sum up to n such that the following conditions are true:
	1. Only numbers 1 through 9 are used.
  2. Each number is used at most once.
	Return a list of all possible valid combinations. The list must not contain the same combination twice,
	and the combinations may be returned in any order.
	Link: https://leetcode.com/problems/combination-sum-iii/
*/

// CombSum find all valid combinations of k numbers that sum up to n
func CombSum(k int, n int) [][]int {
	ans := [][]int{}
	ch := make(chan []int, 1)
	go func(ch chan []int) {
		defer close(ch)
		ns := GenInts(9)
		BackTrack(ch, []int{}, ns, k, n, 0)
	}(ch)
	for comb := range ch {
		ans = append(ans, comb)
	}
	return ans
}

// BackTrack is a helper method for CombSum
func BackTrack(ch chan []int, cand []int, nums []int, k int, n int, pos int) {
	if k == 0 && n == 0 {
		ch <- DeepCopyInts(cand)
	} else {
		for i := pos; i < len(nums); i++ {
			cand = append(cand, nums[i])
			BackTrack(ch, cand, nums, k-1, n-nums[i], i+1)
			cand = cand[:len(cand)-1]
		}
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

// GenInts return an array of int from 1 up to n
func GenInts(n int) []int {
	nums := make([]int, n)
	for i := 1; i < n+1; i++ {
		nums[i-1] = i
	}
	return nums
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
