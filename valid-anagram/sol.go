package sol

/*
	LeetCode Problem 242: Valid Angram
	Level: Easy
	Description: Given two strings s and t, return true if t is an anagram of s, and false otherwise.
	Link: https://leetcode.com/problems/valid-anagram/
*/

// IsAnagram check if two strs are anagram
func IsAnagram(s string, t string) bool {
	return IsTwoMapEqual(CntChFreq(s), CntChFreq(t))
}

// CntChFreq count char frequency of a str
func CntChFreq(str string) map[byte]int {
	freq := map[byte]int{}
	for i := 0; i < len(str); i++ {
		freq[str[i]]++
	}
	return freq
}

// IsTwoMapEqual check if two maps are same
func IsTwoMapEqual(m, n map[byte]int) bool {
	if len(m) != len(n) {
		return false
	}
	for mk, mv := range m {
		nv, ok := n[mk]
		if !ok {
			return false
		}
		if nv != mv {
			return false
		}
	}
	return true
}
