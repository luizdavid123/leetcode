package sol

import "strings"

/*
	LeetCode Problem 6: ZigZag Conversion
	Level: Medium
	Description: Write the code that will take a string and make this conversion given a number of rows.
*/

// ZigZagAndToRow rewrite the string in a zigzag patten and read it one by one by row
func ZigZagAndToRow(str string, n int) string {
	if n == 1 || len(str) == 1 {
		return str
	}
	var builder strings.Builder

	cnt := len(str)
	cyclen := 2 * (n - 1)

	// first row
	for i := 0; i < cnt; i += cyclen {
		builder.WriteByte(str[i])
	}

	// innner row
	for r := 1; r < n-1; r++ {
		for i := r; i < cnt; i += cyclen {
			builder.WriteByte(str[i])
			j := i + 2*(n-r-1)
			if j < cnt {
				builder.WriteByte(str[j])
			}
		}
	}

	// last row
	for i := n - 1; i < cnt; i += cyclen {
		builder.WriteByte(str[i])
	}

	return builder.String()
}
