// 解题思路:位运算
// 时间复杂度O(1), 空间复杂度O(1)
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}

	return false
}