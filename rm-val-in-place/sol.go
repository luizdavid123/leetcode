package sol

/*
	LeetCode Problem 26: Remove Element
	Level: Easy
	Description: Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
	The relative order of the elements may be changed.
*/

// RmEle remove all occurrences of val in nums in-place
func RmEle(nums []int, tar int) ([]int, int) {
	cnt := len(nums)
	if cnt == 0 {
		return []int{}, 0
	}
	k := 0
	for i := 0; i < cnt; i++ {
		if nums[i] != tar {
			nums[k] = nums[i]
			k++
		}
	}
	return nums, k
}
