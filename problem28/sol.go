package problem28

/*
	LeetCode Problem 28: Implement strStr()
	Level: Easy
	Description: Return the index of the first occurrence of needle in haystack,
	or -1 if needle is not part of haystack.
*/

// StrIndxOf return the index of the first occurrence of sub in str, or -1 if not found
func StrIndxOf(str string, sub string) int {
	if len(sub) == 0 {
		return 0
	}
	cntstr, cntsub := len(str), len(sub)
	for i := 0; i < cntstr && cntstr-i >= cntsub; i++ {
		if Match(str[i:], sub) {
			return i
		}
	}
	return -1
}

// Match check if the part include sub
func Match(part string, sub string) bool {
	for i := 0; i < len(sub); i++ {
		if part[i] != sub[i] {
			return false
		}
	}
	return true
}
