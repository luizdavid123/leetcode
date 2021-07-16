package problem135

/*
	LeetCode Problem 135: Candy
	Level: Hard
	Description: There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.
	You are giving candies to these children subjected to the following requirements:
		Each child must have at least one candy.
		Children with a higher rating get more candies than their neighbors.
	Return the minimum number of candies you need to have to distribute the candies to the children.s
	Link: https://leetcode.com/problems/candy/
*/

// MinCandy return the minimum number of candies you need to have to distribute the candies to the children.
func MinCandy(ratings []int) int {
	cnt := len(ratings)
	l2r := make([]int, cnt)
	r2l := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		l2r[i] = 1
		r2l[i] = 1
	}
	for i := 1; i < cnt; i++ {
		if ratings[i] > ratings[i-1] {
			l2r[i] = l2r[i-1] + 1
		}
	}
	for i := cnt - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			r2l[i] = r2l[i+1] + 1
		}
	}
	sum := 0
	for i := 0; i < cnt; i++ {
		sum += MaxInt(l2r[i], r2l[i])
	}
	return sum
}

// MaxInt return the larger one
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
