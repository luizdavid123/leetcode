package sol

import "strings"

/*
	LeetCode Problem 22: Generate Parentheses
	Level: Medium
	Description: Given n pairs of parentheses,
	write a function to generate all combinations of well-formed parentheses.
*/

// GenerateValidParenthesesV2 generate all combinations of well-formed parentheses
func GenerateValidParenthesesV2(n int) []string {
	if n == 0 {
		return []string{""}
	}
	ans := []string{}
	for i := 0; i < n; i++ {
		for _, left := range GenerateValidParenthesesV2(i) {
			for _, right := range GenerateValidParenthesesV2(n - i - 1) {
				ans = append(ans, "("+left+")"+right)
			}
		}
	}
	return ans
}

// GenerateValidParentheses generate all combinations of well-formed parentheses
func GenerateValidParenthesesV3(n int) []string {
	ans := []string{}
	ch := make(chan string)
	go func(ch chan string) {
		defer close(ch)
		BackTrack(ch, "", 0, 0, n)
	}(ch)
	for str := range ch {
		ans = append(ans, str)
	}
	return ans
}

// BackTrack is a helper function to generate parentheses
func BackTrack(ch chan string, str string, open int, close int, max int) {
	if len(str) == 2*max {
		ch <- str
		return
	}
	if close < open {
		BackTrack(ch, str+")", open, close+1, max)
	}
	if open < max {
		BackTrack(ch, str+"(", open+1, close, max)
	}
}

// GenerateValidParentheses generate all combinations of well-formed parentheses
func GenerateValidParentheses(n int) []string {
	valids := []string{}
	pairs := GenPairBracket(n)
	for perm := range GenAllPermStr(pairs) {
		ans := strings.Join(perm, "")
		if Valid(ans) {
			if !Exist(valids, ans) {
				valids = append(valids, ans)
			}
		}
	}
	return valids
}

// Exist check if valid is already in valids
func Exist(valids []string, valid string) bool {
	for i := 0; i < len(valids); i++ {
		if valids[i] == valid {
			return true
		}
	}
	return false
}

// Valid check if the input is valid
func Valid(perm string) bool {
	balance := 0
	for i := 0; i < len(perm); i++ {
		if perm[i] == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

// GenPairBracket generate pairs of brackets
func GenPairBracket(n int) []string {
	pair := make([]string, n*2)
	for i := 0; i < n; i++ {
		pair[2*i], pair[2*i+1] = "(", ")"
	}
	return pair
}

// GenAllPermStr wrap GenPermStr with channel
// Ref: https://github.com/Ramshackle-Jamathon/go-quickPerm
func GenAllPermStr(strs []string) <-chan []string {
	permch := make(chan []string)
	go func(permch chan []string) {
		defer close(permch)
		GenPermStr(permch, strs)
	}(permch)
	return permch
}

// GenPermStr retrun all permutation of the element in strs
// Ref: https://github.com/Ramshackle-Jamathon/go-quickPerm
func GenPermStr(permch chan []string, strs []string) {
	output := make([]string, len(strs))
	copy(output, strs)
	permch <- output

	cnt := len(strs)
	p := make([]int, cnt+1)
	for i := 0; i < cnt+1; i++ {
		p[i] = i
	}

	for i := 1; i < cnt; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		strs[i], strs[j] = strs[j], strs[i]
		output := make([]string, len(strs))
		copy(output, strs)
		permch <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}
