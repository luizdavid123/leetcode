package reversenumber

import "math"

// ReverseNumber return n with its digits reversed.
func ReverseNumber(n int) int {
	ans := 0
	for d := n; d != 0; d = d / 10 {
		ans = ans*10 + d%10
		if ans < math.MinInt32 || ans > math.MaxInt32 {
			return 0
		}
	}
	return ans
}
