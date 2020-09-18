/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{Val:preorder[0]}

    i:=0
    for i < len(inorder) {
        if inorder[i] == root.Val {
            break
        }
        i++
    }
    // 先序计算左子树结束点不好算，根据：中序遍历和先序遍历的左子树相等
    // 先序左子树结尾设为x
    // x - 先序左子树起始 =  中序左子树结束 - 中序左子树起始
    // 代入得：x - 1 = i - 0
    // 所以： x = i + 1 
    root.Left = buildTree(preorder[1:i+1], inorder[:i])
    root.Right = buildTree(preorder[i+1:], inorder[i+1:])
    return root
}