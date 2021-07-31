package sol

/*
	LeetCode Problem 139: Word Break
	Level: Medium
	Description: Given a string s and a dictionary of strings wordDict,
	return true if s can be segmented into a space-separated sequence of one or more dictionary words.
*/

// WordBreak return true if s can be segmented into a space-separated sequence of one or more dictionary words,
// given a string str and a dictionary of strings dict
func WordBreak(str string, dict []string) bool {
	cnt := len(str)
	dp := make([]bool, len(str)+1)
	dp[0] = true

	for i := 1; i <= cnt; i++ {
		for j := 0; j < i; j++ {
			cand := str[j:i]
			if dp[j] && IsStrInStrs(dict, cand) {
				dp[i] = true
				break
			}
		}
	}

	return dp[cnt]
}

// IsStrInStrs check if an element is in an array
func IsStrInStrs(eles []string, ele string) bool {
	cnt := len(eles)
	for i := 0; i < cnt; i++ {
		if eles[i] == ele {
			return true
		}
	}
	return false
}
