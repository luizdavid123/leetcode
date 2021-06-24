package problem17

/*
	LeetCode Problem 17:  Letter Combinations of a Phone Number
	Level: Medium
	Description: Given a string containing digits from 2-9 inclusive, return all possible letter
	combinations that the number could represent. Return the answer in any order.
*/

// AllPhoneNumToLetterComb return all possible letter combinations that the number could represent
func AllPhoneNumToLetterComb(digits string) []string {
	cnt := len(digits)

	cands := make([][]string, cnt)
	for i := 0; i < cnt; i++ {
		cands[i] = GetLetters(digits[i])
	}

	accs := []string{}
	for i := 0; i < len(cands); i++ {
		if i == 0 {
			accs = cands[0]
			continue
		}
		accs = Merge(accs, cands[i])
	}
	return accs
}

// Merge return all combinations of elements in two string slice
func Merge(l, r []string) []string {
	comb := []string{}

	for i := 0; i < len(l); i++ {
		for j := 0; j < len(r); j++ {
			comb = append(comb, l[i]+r[j])
		}
	}

	return comb
}

// GetLetters return corresponding letters with given digit
func GetLetters(digit byte) []string {
	switch digit {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	default:
		return nil
	}
}
