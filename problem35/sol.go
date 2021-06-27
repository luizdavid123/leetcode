package problem35

/*
	LeetCode Problem 35: Search Insert Position
	Level: Medium
	Description: Given a sorted array of distinct integers and a target value,
	return the index if the target is found.
	If not, return the index where it would be if it were inserted in order.
*/

// SearchOrInsertIfNotFound return the index if the target is found or return
// the index where it would be if it were inserted in order.
func SearchOrInsertIfNotFound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r + l) / 2
		v := nums[m]
		switch {
		case target == v:
			return m
		case target > v:
			l = m + 1
		case target < v:
			r = m - 1
		}
	}
	return l
}
