/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// terminal:没找到或者找到了其中一个节点:p,q
	if root == nil || root == p || root == q {
		return root
	}
	// drop left right
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

   // 左边空，返回右边
   if l == nil {
	   return r
   }
   // 右边空，返回左边
   if r == nil {
	   return l
   }
   // 都不空, 返回当前节点，说明找到了
	return root
}