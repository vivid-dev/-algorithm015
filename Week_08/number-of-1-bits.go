// 解题思路:位运算
// 时间复杂度O(N), 空间复杂度O(1)
func hammingWeight(num uint32) int {
	cnt := 0
	for num > 0 {
		// 判断末尾是否为1
		if num&1 > 0 {
			cnt++
		}
		num >>= 1
	}

	return cnt
}