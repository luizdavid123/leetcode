package problem202

/*
	LeetCode Problem 202: Happy Number
	Level: Easy
	Description: Write an algorithm to determine if a number n is happy.
	Link: https://leetcode.com/problems/happy-number/
*/

// IsHappy determine if a number n is happy
func IsHappy(n int) bool {
	fast, slow := n, n
	for {
		slow = DigitSquereSum(slow)
		fast = DigitSquereSum(fast)
		fast = DigitSquereSum(fast)
		if fast == 1 {
			return true
		}
		if fast == slow {
			return false
		}
	}
}

// DigitSquereSum reutrn the sum of the squares of num's digits
func DigitSquereSum(n int) int {
	sum := 0
	for m := n; m > 0; m /= 10 {
		d := m % 10
		sum += d * d
	}
	return sum
}
