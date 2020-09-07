// BFS
// 解题思路：使用队列做, golang没有队列使用切片模拟队列操作
// cnt 用来记录当前队列长度
// i 遍历当前队列内容
// j 遍历孩子节点
// 复杂度: 时间复杂度O(N), 空间复杂度O(1)

func levelOrder(root *Node) [][]int {
    res := make([][]int, 0)
    if root == nil {
        return res
    }

    queue := make([]*Node, 0)
    queue = append(queue, root)

    for len(queue) > 0 {
        cnt := len(queue)
        ans := make([]int, 0)
        // 处理一层
        for i := 0; i < cnt; i++ {
            // 孩子节点入队
            for j := 0; j < len(queue[0].Children); j++ {
                queue = append(queue, queue[0].Children[j])
            }
            // 结果输出到层队列里
            ans = append(ans, queue[0].Val)
            queue = queue[1:]
        }

        res = append(res, ans)
    }

    return res
}