// 循环迭代解法
// 思路: 由于数组长度固定 & 每次挪到很多数据
// 从后向前来合并
// i 迭代nums1
// j 迭代nums2
// cur 标识要填的位置
func merge(nums1 []int, m int, nums2 []int, n int)  {
    // 迭代 nums1 & nums2
    i := m - 1
    j := n - 1
    cur := len(nums1) - 1
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[cur] = nums1[i]
            i--
        } else {
            nums1[cur] = nums2[j]
            j--
        }
        cur--
    }
    // nums2没跑完继续迭代nums2
    for j >= 0 {
        nums1[cur] = nums2[j]
        cur--
        j--
    }
}
