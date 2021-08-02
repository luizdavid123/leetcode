package sol

/*
	LeetCode Problem 229: Majority Element II
	Level: Easy
	Description: Given an array nums of size n, return the majority element.
	The majority element is the element that appears more than ⌊n / 3⌋ times. You may assume that the majority element always exists in the array.
	Link: https://leetcode.com/problems/majority-element/
*/

// FindMajorEle find the majority element that appears more than ⌊n / 2⌋ times
func FindMajorEle(nums []int) []int {
	cand_one, cand_two := nums[0], nums[0]
	cnt_one, cnt_two := 0, 0

	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] == cand_one:
			cnt_one++
		case nums[i] == cand_two:
			cnt_two++
		case cnt_one == 0:
			cand_one, cnt_one = nums[i], 1
		case cnt_two == 0:
			cand_two, cnt_two = nums[i], 1
		default:
			cnt_one--
			cnt_two--
		}
	}

	cnt_one, cnt_two = 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == cand_one {
			cnt_one++
		} else if nums[i] == cand_two {
			cnt_two++
		}
	}
	ans := []int{}
	if cnt_one > len(nums)/3 {
		ans = append(ans, cand_one)
	}
	if cnt_two > len(nums)/3 {
		ans = append(ans, cand_two)
	}
	return ans
}
