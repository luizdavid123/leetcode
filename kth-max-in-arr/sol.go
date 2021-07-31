package sol

import "sort"

/*
	LeetCode Problem 215: Kth Largest Element in an Array
	Level: Medium
	Description: Given an integer array nums and an integer k, return the kth largest element in the array.
	Link: https://leetcode.com/problems/kth-largest-element-in-an-array/
*/

// FindKthTop return the kth largest element in the array.
func FindKthTop(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// FindKthTop return the kth largest element in the array.
func FindKthTopV2(nums []int, k int) int {
	k = len(nums) - k
	high, low := 0, len(nums)-1
	for low < high {
		pivot := Partition(nums, high, low)
		if pivot < k {
			low = pivot + 1
		} else if pivot > k {
			high = pivot - 1
		} else {
			break
		}
	}
	return nums[k]
}

// Partition divide the arr into left half and right half, nums in the left half <= pivot,
// nums in the right half > pivot
func Partition(nums []int, high int, low int) int {
	pivot := nums[high]
	i := low
	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[high], nums[i] = nums[i], nums[high]
	return i
}
