// 解题思路：DP
// 时间复杂度O(N), 空间复杂度O(1)
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }

    return max(helper(nums[1:]), helper(nums[:len(nums)-1]))
}


func helper(nums []int) int {
    pre, now := 0, 0
    for i := 0; i < len(nums); i++ {
        pre, now = now, max(now, pre+nums[i])
    }

    return now
}

func max(x, y int) int {
    if x > y {
        return x
    }

    return y
}