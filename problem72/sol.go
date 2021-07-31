package sol

/*
	LeetCode Problem 72: Edit Distance
	Level: Hard
	Description: Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
	You have the following three operations permitted on a word:
		1. Insert a character
		2. Delete a character
		3. Replace a character
*/

// MinEditDist return the minimum number of operations required to convert lhs to rhs
func MinEditDist(lhs, rhs string) int {
	n, m := len(lhs), len(rhs)
	cost := make([][]int, n+1)
	for i := range cost {
		cost[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		cost[i][0] = i
	}
	for i := 1; i <= m; i++ {
		cost[0][i] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if lhs[i] == rhs[j] {
				cost[i+1][j+1] = cost[i][j]
			} else {
				replace := cost[i][j]
				delete := cost[i][j+1]
				insert := cost[i+1][j]
				cost[i+1][j+1] = MinInt(replace, MinInt(delete, insert)) + 1
			}
		}
	}

	return cost[n][m]
}

// MinInt return the smaller int
func MinInt(n, m int) int {
	if n > m {
		return m
	}
	return n
}
