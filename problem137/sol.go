package problem137

/*
	LeetCode Problem 137: Single Number II
	Level: Medium
	Description: Given an integer array nums where every element appears three times except for one,
	which appears exactly once. Find the single element and return it.
*/

// FindSingleNum find the number that only appears once
func FindSingleNum(nums []int) int {
	cnt := len(nums)
	counter := make(map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		count, ok := counter[nums[i]]
		if !ok {
			counter[nums[i]] = 1
		} else {
			counter[nums[i]] = count + 1
		}
	}

	for num, count := range counter {
		if count == 1 {
			return num
		}
	}
	return -1
}
