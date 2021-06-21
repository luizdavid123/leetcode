package problem5

/*
	LeetCode Problem 5: Longest Palindromic Substring
	Level: Medium
	Description: Given a string s, return the longest palindromic substring in s.
*/

// LongestParaSubStr return the longest palindromic substring
func LongestParaSubStr(str string) string {
	odd := LongestOddParaSubStr(str)
	even := LongestEvenParaSubStr(str)
	if len(odd) > len(even) {
		return odd
	}
	return even
}

// LongestParaSubStrV2 return the longest palindromic substring using another algorithm
func LongestParaSubStrV2(str string) string {
	return ""
}

// IsStrPara check if the str is palindormic
func IsStrPara(str string) bool {
	if len(str)%2 == 1 {
		return IsOddStrPara(str)
	}
	return IsEvenStrPara(str)
}

// IsOddStrPara check if the str with odd length is palindormic
func IsOddStrPara(str string) bool {
	if len(str) == 1 {
		return true
	}
	m := len(str) / 2
	for l, r := m-1, m+1; l >= 0 && r < len(str); l, r = l-1, r+1 {
		if str[l] != str[r] {
			return false
		}
	}
	return true
}

// IsEvenPara check if the str with even length is palindormic
func IsEvenStrPara(str string) bool {
	m := len(str)/2 - 1
	for l, r := m, m+1; l >= 0 && r < len(str); l, r = l-1, r+1 {
		if str[l] != str[r] {
			return false
		}
	}
	return true
}

// LongestOddParaSubStr return the longest palindromic substring from origin odd length string
func LongestOddParaSubStr(str string) string {
	res := str[0:1]
	max := 0
	for i := 0; i < len(str); i++ {
		for l, r := i-1, i+1; l >= 0 && r < len(str); l, r = l-1, r+1 {
			if str[l] != str[r] {
				break
			}
			if max < r-l+1 {
				res = str[l : r+1]
				max = r - l + 1
			}
		}
	}
	return res
}

// LongestEvenParaSubStr return the longest palindromic substring from origin even length string
func LongestEvenParaSubStr(str string) string {
	res := str[0:1]
	max := 0
	for i := 0; i < len(str); i++ {
		for l, r := i, i+1; l >= 0 && r < len(str); l, r = l-1, r+1 {
			if str[l] != str[r] {
				break
			}
			if max < r-l+1 {
				res = str[l : r+1]
				max = r - l + 1
			}
		}
	}
	return res
}
