// 解题思路：DP
// 时间复杂度O(N), 空间复杂度O(1)
func rob(nums []int) int {
    pre, now := 0, 0

    for i := 0; i < len(nums); i++ {
        // dp[i] = max(dp[i - 1], dp[i-2] + nums[i])
        pre, now = now, max(now, pre + nums[i])
    }

    return now
}

func max(x, y int) int {
    if x > y {
        return x
    }

    return y
}