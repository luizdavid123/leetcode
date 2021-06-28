package problem38

import "fmt"

/*
	LeetCode Problem 38: Count and Say
	Level: Medium
	Description: Given a positive integer n, return the nth term of the count-and-say sequence.
*/

// CountAndSay return return the nth term of the count-and-say sequence.
func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	latest := CountAndSay(n - 1)

	ans := ""
	ch := latest[0]
	cnt := 1
	for i := 1; i < len(latest); i++ {
		if ch != latest[i] {
			ans += fmt.Sprintf("%d%s", cnt, string(ch))
			ch = latest[i]
			cnt = 1
		} else {
			cnt++
		}
	}
	ans += fmt.Sprintf("%d%s", cnt, string(ch))
	return ans
}
