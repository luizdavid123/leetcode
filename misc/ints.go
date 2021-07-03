package misc

// MaxInt return the larger int
func MaxInt(n, m int) int {
	if n > m {
		return n
	}
	return m
}

// MinInt return the smaller int
func MinInt(n, m int) int {
	if n > m {
		return m
	}
	return n
}

// AbsInt return the absolute value of n
func AbsInt(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

// IsIntInInts check if the int is in ints
func IsIntInInts(nums []int, num int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return true
		}
	}
	return false
}

// IndexInt return the index of first ocurrence of n or return -1 if not found
func IndexInt(nums []int, num int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return i
		}
	}
	return -1
}

// RvsIns reverse a list of int
func RvsIns(nums []int) []int {
	cnt := len(nums)
	for i := 0; i < cnt/2; i++ {
		nums[i], nums[cnt-1-i] = nums[cnt-1-i], nums[i]
	}
	return nums
}
