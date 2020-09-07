// dfs解法
// 思路: 树的先序遍历，使用栈实现
// 使用数组实现stack
// 遍历结果存储到res里 
type stack struct {
	Stack []*TreeNode
	Len int
}


func (s *stack) push(node *TreeNode) {
	s.Stack = append(s.Stack, node)
	s.Len++
}

func (s *stack) pop() (node *TreeNode) {
	node = s.Stack[s.Len-1]
	s.Stack = s.Stack[0:s.Len-1]
	s.Len--
	return 
}

func newStack() *stack{
    return &stack{Stack:[]*TreeNode{}, Len:0}
}

func preorderTraversal(root *TreeNode) []int {
	s:=newStack()
	s.push(root)
	res := make([]int, 0)
	for s.Len > 0 {
		tmp := s.pop()
		if tmp != nil {
			res = append(res, tmp.Val)
            if tmp.Right != nil {
				s.push(tmp.Right)
			}
			if tmp.Left != nil {
				s.push(tmp.Left)
			}	
		}
	}
    return res
}