package sol

/*
	LeetCode Problem 81: Combinations
	Level: Medium
	Description: Given two integers n and k,
	return all possible combinations of k numbers out of the range [1, n].
*/

// GenComb return all combs of k numbers out of the range [1, n]
func GenComb(n int, k int) [][]int {
	ans := [][]int{}
	ch := make(chan []int, 1)
	go func(ch chan []int) {
		defer close(ch)
		BackTrack(ch, []int{}, n, k, 1)
	}(ch)
	for comb := range ch {
		ans = append(ans, comb)
	}
	return ans
}

// BackTrack is a helper method for GenComb
func BackTrack(ch chan []int, cand []int, upper int, limit int, pos int) {
	if limit == 0 {
		ch <- DeepCopyInts(cand)
	} else {
		for i := pos; i <= upper; i++ {
			cand = append(cand, i)
			BackTrack(ch, cand, upper, limit-1, i+1)
			cand = cand[:len(cand)-1]
		}
	}
}

// GenCombV2 return all combs of k numbers out of the range [1, n]
func GenCombV2(n int, k int) [][]int {
	ans := [][]int{}
	ch := make(chan []int, 1)
	go func(ch chan []int) {
		defer close(ch)
		ns := GenInts(n)
		BackTrackV2(ch, []int{}, ns, k, 0)
	}(ch)
	for comb := range ch {
		ans = append(ans, comb)
	}
	return ans
}

// BackTrackV2 is a helper method for GenComb
func BackTrackV2(ch chan []int, cand []int, nums []int, limit int, pos int) {
	if limit == 0 {
		ch <- DeepCopyInts(cand)
	} else {
		for i := pos; i < len(nums); i++ {
			cand = append(cand, nums[i])
			BackTrackV2(ch, cand, nums, limit-1, i+1)
			cand = cand[:len(cand)-1]
		}
	}
}

// DeepCopyInts return a copy of int slice
func DeepCopyInts(nums []int) []int {
	copy := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		copy[i] = nums[i]
	}
	return copy
}

// GenInts return an array of int from 1 up to n
func GenInts(n int) []int {
	nums := make([]int, n)
	for i := 1; i < n+1; i++ {
		nums[i-1] = i
	}
	return nums
}

// IsIntInInts check if int is in ins
func IsIntInInts(nums []int, num int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return true
		}
	}
	return false
}
