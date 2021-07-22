package problem167

/*
	LeetCode Problem 167: Two Sum II - Input array is sorted
	Level: Easy
	Description: Given an array of integers numbers that is already sorted in non-decreasing order,
	find two numbers such that they add up to a specific target number.
	Link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
*/

// TwoSum	find two numbers such that they add up to tar
func TwoSum(nums []int, tar int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		comp := tar - nums[i]
		j := BinearySearch(nums, comp, 0, len(nums)-1)
		if j != -1 && j != i {
			if j > i {
				return []int{i, j}
			}
			return []int{j, i}
		}
	}
	return ans
}

// BinearySearch return the index of tar if found, othervise return -1
func BinearySearch(nums []int, tar int, left int, right int) int {
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == tar {
			return mid
		}
		if nums[mid] > tar {
			right = mid
		}
		if nums[mid] < tar {
			left = mid + 1
		}
	}
	return -1
}
