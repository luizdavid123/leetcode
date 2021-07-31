package sol

/*
	LeetCode Problem 1: Two Sum
	Level: Easy
	Description: Given an array of integers nums and an integer target,
	return indices of the two numbers such that they add up to target.
*/

// BruteForce return indices of the two numbers such that they add up to target.
func BruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if i != j && sum == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// HashMap return indices of the two numbers such that they add up to target.
func HashMap(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		j, ok := m[target-v]
		if ok && i != j {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}
