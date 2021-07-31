package sol

/*
	LeetCode Problem 91: Decode Ways
	Level: Medium
	Description: Given a string s containing only digits, return the number of ways to decode it.
*/

// NumOfDecodeWay return the number of ways to decode a string containing only digits
func NumOfDecodeWay(str string) int {
	if str[0] == '0' {
		return 0
	}
	cnt := len(str)
	dp := make([]int, cnt+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= cnt; i++ {
		if ValidCode(str[i-1 : i]) {
			dp[i] += dp[i-1]
		}
		if ValidCode(str[i-2 : i]) {
			dp[i] += dp[i-2]
		}
	}
	return dp[cnt]
}

// ValidCode check if the group is a valid code
func ValidCode(group string) bool {
	cnt := len(group)
	if cnt > 3 || cnt == 0 {
		return false
	}
	if group[0] == '0' {
		return false
	}
	switch {
	case cnt == 1:
		return true
	case cnt == 2:
		if group[0] > '2' {
			return false
		}
		if group[0] == '2' && group[1] > '6' {
			return false
		}
	}
	return true
}
