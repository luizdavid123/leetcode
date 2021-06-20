package maxarea

/*
	LeetCode Problem 11 Palindrome Number
	Level: Medium
	Description: Given n non-negative integers a1, a2, ..., an ,
	where each represents a point at coordinate (i, ai).
	n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0).
	Find two lines, which, together with the x-axis forms a container,
	such that the container contains the most water.
*/

// MaxArea find two lines together with the x-axis forms a container such that
// the container contains the most water.
func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		ans = Max(ans, Min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

// Min return the smaller value between two numbers
func Min(n, m int) int {
	if n < m {
		return n
	}
	return m
}

// Max return the larger value between two numbers
func Max(n, m int) int {
	if n < m {
		return m
	}
	return n
}
