package problem169

/*
	LeetCode Problem 169: Majority Element
	Level: Easy
	Description: Given an array nums of size n, return the majority element.
	The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
	Link: https://leetcode.com/problems/majority-element/
*/

// FindMajorEle find the majority element that appears more than ⌊n / 2⌋ times
func FindMajorEle(nums []int) int {
	counter := map[int]int{}
	for i := 0; i < len(nums); i++ {
		count, ok := counter[nums[i]]
		if !ok {
			counter[nums[i]] = 1
		} else {
			counter[nums[i]] = count + 1
		}
	}
	for num, cnt := range counter {
		if cnt > len(nums)/2 {
			return num
		}
	}
	return -1
}
