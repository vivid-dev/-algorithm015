// 递归解法
// 思路: 递归遍历两个链表
// 由于要升序，所以选择较小值挂链
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val > l2.Val {
        l2.Next = mergeTwoLists(l1, l2.Next)
        return l2
    }
    l1.Next = mergeTwoLists(l1.Next, l2)
    
    return l1
}
