package misc

// DeepCopyInts return a copy of int slice
func DeepCopyInts(nums []int) []int {
	copy := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		copy[i] = nums[i]
	}
	return copy
}

// DeepCopyStrs return a copy of string slice
func DeepCopyStrs(strs []string) []string {
	copy := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		copy[i] = strs[i]
	}
	return copy
}
