// dfs
// 思路：模拟二叉树先序:根、左、右的顺序,先存入root，
// 再从左至右存入孩子节点
// res 存储结果
// 复杂度: 时间复杂度O(N), 空间复杂度O(N)
func preorder(root *Node) []int {
    res := make([]int, 0)
    if root == nil {
        return res
	}
	
	res = append(res, root.Val)
	for i := range root.Children {
		res = append(res, preorder(root.Children[i])...)
	}
	return res
}