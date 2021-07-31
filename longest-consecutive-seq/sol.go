package sol

/*
	LeetCode Problem 128: Longest Consecutive Sequence
	Level: Medium
	Description: Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
	Link: https://leetcode.com/problems/longest-consecutive-sequence/
*/

// LongestConsective return the length of the longest consecutive elements sequence.
func LongestConsective(nums []int) int {
	set := map[int]int{}
	for _, num := range nums {
		cnt, ok := set[num]
		if !ok {
			set[num] = 0
		}
		set[num] = cnt + 1
	}
	ans := 0
	for num := range set {
		if _, ok := set[num-1]; !ok {
			curr := num
			streak := 1
			for {
				if _, ok := set[curr+1]; ok {
					curr += 1
					streak += 1
					continue
				} else {
					break
				}
			}
			ans = MaxInt(ans, streak)
		}
	}
	return ans
}

// MaxInt return the larger one
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
