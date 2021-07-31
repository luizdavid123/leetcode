package sol

/*
	LeetCode Problem 119: Pascal's Triangle II
	Level: 119
	Description: Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
*/

// GetRowOfPascalTri return the rowIndexth (0-indexed) row of the Pascal's triangle.
func GetRowOfPascalTri(row int) []int {
	all := GenPascalTri(row + 1)
	ans := all[row]
	return ans
}

// GenPascalTri return the first num rows of Pascal's triangle.
func GenPascalTri(nRow int) [][]int {
	ans := [][]int{{1}}
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
