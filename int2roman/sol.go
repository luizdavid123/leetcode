package int2roman

/*
	LeetCode Proble12: "m", Integer to Roman
	Level: Medium
	Description: Given an integer, convert it to a roman numeral.
*/

// IntToRoman convert an integer to a roman numeral
func IntToRoman(n int) string {
	v2s := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	vs := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	ans := ""
	r := n
	for _, v := range vs {
		m := r / v
		for i := 0; i < m; i++ {
			ans += v2s[v]
		}
		r -= v * m
	}

	return ans
}
