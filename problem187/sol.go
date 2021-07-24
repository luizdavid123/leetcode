package problem187

/*
	LeetCode Problem 187: Repeated DNA Sequences
	Level: Medium
	Description: Given a string s that represents a DNA sequence,
	return all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.
	You may return the answer in any order.
	Link: https://leetcode.com/problems/repeated-dna-sequences/
*/

// FindRepeatedDNASeq all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.
func FindRepeatedDNASeq(str string) []string {
	seen, repeat := map[string]int{}, map[string]int{}
	for i := 0; i+9 < len(str); i++ {
		sub := str[i : i+10]
		cnt, ok := seen[sub]
		if !ok {
			seen[sub] = 1
		} else {
			seen[sub] = cnt + 1
			repeat[sub] = 1
		}
	}
	ans := []string{}
	for sub := range repeat {
		ans = append(ans, sub)
	}
	return ans
}
