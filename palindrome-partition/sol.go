package sol

/*
	LeetCode Problem 131: Palindrome Partitioning
	Level: Medium
	Description: Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.
	Link: https://leetcode.com/problems/palindrome-partitioning/
*/

// Partition return all possible partitioning of str such that
// every substring of the  partition is a palindrome.
func Partition(str string) [][]string {
	ans := [][]string{}
	ch := make(chan []string)
	go func(ch chan []string) {
		defer close(ch)
		BackTrack(ch, str, []string{}, 0)
	}(ch)
	for cand := range ch {
		ans = append(ans, cand)
	}
	return ans
}

// BackTrack is the helper method for Partition
func BackTrack(ch chan []string, str string, cand []string, start int) {
	if start >= len(str) {
		ch <- DeepCopyStrs(cand)
		return
	}
	for i := start; i < len(str); i++ {
		tmp := str[start : i+1]
		if IsPalindrome(tmp) {
			cand = append(cand, tmp)
			BackTrack(ch, str, cand, i+1)
			cand = cand[:len(cand)-1]
		}
	}
}

// DeepCopyStrs return a copy of string slice
func DeepCopyStrs(strs []string) []string {
	copy := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		copy[i] = strs[i]
	}
	return copy
}

// IsPalindrome check if it is a palindrome
func IsPalindrome(str string) bool {
	cand := str
	cnt := len(cand)
	for i := 0; i < cnt/2; i++ {
		if cand[i] != cand[cnt-i-1] {
			return false
		}
	}
	return true
}
