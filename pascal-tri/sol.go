package sol

/*
	LeetCode Problem 118:  Pascal's Triangle
	Level: Easy
	Description: Given an integer numRows, return the first num rows of Pascal's triangle.
*/

// GenPascalTri return the first num rows of Pascal's triangle.
func GenPascalTri(nRow int) [][]int {
	ans := [][]int{}

	ans = append(ans, []int{1})

	for i := 1; i < nRow; i++ {
		prev := ans[i-1]

		row := []int{1}
		for j := 1; j < len(prev); j++ {
			row = append(row, prev[j-1]+prev[j])
		}
		row = append(row, 1)

		ans = append(ans, row)
	}

	return ans
}
