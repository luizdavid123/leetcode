package problem3

/*
	LeetCode Problem 3: Longest Substring Without Repeating Characters
	Level: Medium
	Description: Given a string s, find the length of the longest substring without repeating characters.
*/

// LongestSubStr find the length of the longest substring without repeating characters
func LongestSubStr(str string) int {
	max := 0
	visited := make(map[byte]int)

	for start, end := 0, 0; end < len(str); end++ {
		if _, ok := visited[str[end]]; ok {
			start = MaxInt(start, visited[str[end]])
		}
		max = MaxInt(max, end-start+1)
		visited[str[end]] = end + 1
	}

	return max
}

// MaxInt return the larger interer
func MaxInt(left, right int) int {
	if left <= right {
		return right
	}
	return left
}
