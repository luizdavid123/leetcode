package sol

/*
	LeetCode Problem 20: Valid Parentheses
	Level: Easy
	Description: Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
	An input string is valid if:
	1. Open brackets must be closed by the same type of brackets.
	2. Open brackets must be closed in the correct order.
*/

// ValidParentheses check if the input is valid
func ValidParentheses(input string) bool {
	c2o := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	cnt := len(input)

	stack := []byte{}
	for i := 0; i < cnt; i++ {
		ch := input[i]
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			size := len(stack)
			if size == 0 {
				return false
			}
			top := stack[size-1]
			if top != c2o[ch] {
				return false
			}
			stack = stack[:size-1]
		}
	}

	return len(stack) == 0
}
