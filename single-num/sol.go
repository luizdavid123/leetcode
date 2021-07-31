package sol

/*
	LeetCode Problem 136: Single Number
	Level: Easy
	Description: Given a non-empty array of integers nums, every element appears twice except for one.
	Find that single one.
*/

/*
	Bit Manipulation Tips:
	0. 0 xor 0 = 0, 1 xor 0 = 1, 0 xor 1 = 1, 1 xor 1 = 0
	1. a xor 0 = a
	2. a xor a = 0
	so, a xor b xor a = a xor a xor b = 0 xor b = b
*/

// FindSingleNum find the number that only appears once
func FindSingleNum(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}
