// 解题思路：DP
// 时间复杂度O(N), 空间复杂度O(N)
func massage(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])

    for i := 2; i < len(dp); i++ {
        dp[i] = max(dp[i-1], dp[i-2] + nums[i])
    }

    return dp[len(dp)-1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}