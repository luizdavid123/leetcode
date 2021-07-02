package problem49

/*
	LeetCode Problem 49: Group Anagrams
	Level: Medium
	Description: Given an array of strings strs, group the anagrams together.
	You can return the answer in any order.
	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
	typically using all the original letters exactly once.
*/

// GroupAnagrams group the anagrams together
func GroupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	for i := 0; i < len(strs); i++ {
		if len(ans) == 0 {
			ans = append(ans, []string{strs[i]})
		} else {
			idx := IndexOf(ans, strs[i])
			if idx == -1 {
				ans = append(ans, []string{strs[i]})
			} else {
				ans[idx] = append(ans[idx], strs[i])
			}
		}
	}
	return ans
}

func IndexOf(group [][]string, str string) int {
	cnt := len(group)
	for i := 0; i < cnt; i++ {
		cand := group[i][0]
		if IsTwoStrAnag(cand, str) {
			return i
		}
	}
	return -1
}

// IsTwoStrAnag check if a string is an anagram to other string
func IsTwoStrAnag(lhs, rhs string) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	l2c := map[byte]int{}
	for i := 0; i < len(lhs); i++ {
		ch := lhs[i]
		cnt, ok := l2c[ch]
		if !ok {
			l2c[ch] = 1
		} else {
			l2c[ch] = cnt + 1
		}
	}
	for i := 0; i < len(rhs); i++ {
		ch := rhs[i]
		cnt, ok := l2c[ch]
		if !ok {
			return false
		}
		cnt--
		if cnt < 0 {
			return false
		}
		l2c[ch] = cnt
	}
	return true
}
