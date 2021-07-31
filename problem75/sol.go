package sol

/*
	LeetCode Problem 75: Sort Colors
	Level: Medium
	Description: Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent,
	with the colors in the order red, white, and blue. We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
	You must solve this problem without using the library's sort function.
*/

// SortColors sort an array of colors
func SortColors(nums []int) []int {
	counter := [3]int{0, 0, 0}

	for _, num := range nums {
		counter[num]++
	}

	index := 0
	for num, count := range counter {
		for i := 0; i < count; i++ {
			nums[index] = num
			index++
		}
	}

	return nums
}
