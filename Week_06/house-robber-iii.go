// 解题思路:dp
// 时间复杂度O(N), 空间复杂度O(1)
func rob(root *TreeNode) int {
    res := helper(root)
    return max(res[0], res[1])
}

func helper(root *TreeNode) []int {
    if root == nil {
        return []int{0, 0}
    }
    left := helper(root.Left)
    right := helper(root.Right)

    dp := make([]int, 2)
    dp[0] = max(left[0], left[1]) + max(right[0], right[1])
    dp[1] = root.Val + left[0] + right[0]

    return dp
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}