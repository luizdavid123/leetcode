package sol

/*
	LeetCode Problem 47: Permutations II
	Level: Medium
	Description: Given a collection of numbers, nums, that might contain duplicates,
	return all possible unique permutations in any order.
*/

/*
	As a reminder, backtracking is a general algorithm for finding all (or some) solutions to some problems
	with constraints. It incrementally builds candidates to the solutions, and abandons a candidate as soon
	as it determines that the candidate cannot possibly lead to a solution.
*/

/*
	Ref: https://leetcode.com/problems/permutations-ii/solution/
*/

// UniquePerm return all possible unique permutations in any order
func UniquePerm(nums []int) [][]int {
	ans := [][]int{}
	ch := make(chan []int)
	go func(ch chan []int) {
		defer close(ch)
		counter := map[int]int{}
		for _, num := range nums {
			cnt, ok := counter[num]
			if !ok {
				counter[num] = 0
			}
			counter[num] = cnt + 1
		}
		comb := []int{}
		BackTrack(ch, comb, len(nums), counter)
	}(ch)
	for comb := range ch {
		ans = append(ans, comb)
	}
	return ans
}

// BackTrack is the helper method for finding all unique permutation
func BackTrack(ch chan []int, comb []int, cnt int, counter map[int]int) {
	if len(comb) == cnt {
		copy := make([]int, len(comb))
		for i := 0; i < len(comb); i++ {
			copy[i] = comb[i]
		}
		ch <- copy
	}
	for num, count := range counter {
		if count == 0 {
			continue
		}
		comb = append(comb, num)
		counter[num]--
		BackTrack(ch, comb, cnt, counter)
		comb = comb[:len(comb)-1]
		counter[num] = count
	}
}
