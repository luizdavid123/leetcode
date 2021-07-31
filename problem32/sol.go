package sol

/*
	LeetCode Problem 32: Longest Valid Parentheses
	Level: Hard
	Description: Given a string containing just the characters '(' and ')',
	find the length of the longest valid (well-formed) parentheses substring.
*/

// LongestValidParentheses find the length of the longest valid (well-formed) parentheses substring.
func LongestValidParentheses(str string) int {
	ans := 0
	stack := []int{-1}
	top := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			stack = append(stack, i)
			top++
		} else {
			stack = stack[:top]
			top--
			if len(stack) == 0 {
				stack = append(stack, i)
				top++
			} else {
				v := stack[top]
				ans = MaxInt(ans, i-v)
			}
		}
	}
	return ans
}

// MaxInt return the larger int
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
