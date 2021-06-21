package problem13

/*
	LeetCode Problem 13: Roman to Integer
	Level: Easy
	Description: Given a roman numeral, convert it to an integer.
*/

// RomToInt convert a roman numeral to an integer
func RomToInt(str string) int {
	s2v := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	num := 0
	cnt := len(str)
	for i := 0; i < cnt-1; i++ {
		curr := s2v[str[i]]
		next := s2v[str[i+1]]
		if next > curr {
			num -= s2v[str[i]]
		} else {
			num += s2v[str[i]]
		}
	}
	num += s2v[str[cnt-1]]
	return num
}
