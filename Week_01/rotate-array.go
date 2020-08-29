// 线性代数矩阵求逆秩
// 思路：分别对前半段、后半段、整体求逆秩
func rotate(nums []int, k int)  {
    idx := len(nums) - k % len(nums)
    if idx < 0 {
        return
    }
    reverse(nums, 0, idx - 1)
    reverse(nums, idx, len(nums) - 1)
    reverse(nums, 0, len(nums) - 1)
}

func reverse(nums []int, l, r int) {
    for i, j := l, r; i < j; {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}
