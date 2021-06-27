package problem30

/*
	LeetCode Problem 30: Substring with Concatenation of All Words
	Level: Hard
	Description: You are given a string s and an array of strings words of the same length.
	Return all starting indices of substring(s) in s that is a concatenation of each word
	in words exactly once, in any order, and without any intervening characters.
*/

// FindSubStr return all starting indices of substring(s) in s that is a concatenation of each word
// in words exactly once, in any order, and without any intervening characters
func FindSubStr(str string, words []string) []int {
	counts := map[string]int{}
	for i := 0; i < len(words); i++ {
		count, ok := counts[words[i]]
		if !ok {
			counts[words[i]] = 1
		} else {
			counts[words[i]] = count + 1
		}
	}

	idx := []int{}
	strlen, nwords, wordlen := len(str), len(words), len(words[0])
	for i := 0; i < strlen-nwords*wordlen+1; i++ {
		sub := str[i : i+nwords*wordlen]
		if IsSubConCatOfWord(sub, counts, wordlen) {
			idx = append(idx, i)
		}
	}
	return idx
}

// IsSubConCatOfWord check if the sub contains the concat of words
func IsSubConCatOfWord(sub string, counts map[string]int, wordlen int) bool {
	seen := map[string]int{}
	for i := 0; i < len(sub); i += wordlen {
		word := sub[i : i+wordlen]
		count, ok := seen[word]
		if !ok {
			seen[word] = 1
		} else {
			seen[word] = count + 1
		}
	}
	return IsTwoStrIntMapEqual(counts, seen)
}

// IsTwoMapEqual check if two maps are equal. Two maps are equal if their length is the same
// and all key-value pairs are the same
func IsTwoStrIntMapEqual(left, right map[string]int) bool {
	if len(left) != len(right) {
		return false
	}
	for lk, lv := range left {
		rv, ok := right[lk]
		if !ok {
			return false
		}
		if rv != lv {
			return false
		}
	}
	return true
}
