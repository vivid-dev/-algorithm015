// dfs解法
// 思路: 根据二叉树的中序遍历，将遍历的结果保存并返回
// ans 存放遍历的结果
// 复杂度: 时间复杂度O(N), 空间复杂度O(N)
func inorderTraversal(root *TreeNode) []int {
    if root != nil {
        ans := inorderTraversal(root.Left)
        ans = append(ans, root.Val)
        ans = append(ans, inorderTraversal(root.Right)...)
        return ans
    }
    return []int{}
}