// 解题思路：并查集
// 时间复杂度O(2^N), 空间复杂度O(N)
func helper(left, right int, ans string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, ans)
		return
	}

	if left > 0 {
		helper(left-1, right, ans+"(", res)
	}

	if right > left {
		helper(left, right-1, ans+")", res)
	}
}

func generateParenthesis(n int) []string {
	res := []string{}
	helper(n, n, "", &res)
	return res
}