// 解题思路：DP
// 时间复杂度O(N), 空间复杂度O(N)
func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return math.MinInt32
    }
    dp := make([][]int, len(nums))
    for i := range dp {
        dp[i] = make([]int, 2)
    }
    dp[0][0] = nums[0]
    dp[0][1] = nums[0]
    max := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] > 0 {
            dp[i][0] = Max(nums[i], dp[i-1][0]*nums[i])
            dp[i][1] = Min(nums[i], dp[i-1][1]*nums[i])
        } else {
            dp[i][0] = Max(nums[i], dp[i-1][1]*nums[i])
            dp[i][1] = Min(nums[i], dp[i-1][0]*nums[i])
        }
        max = Max(max, dp[i][0])
    }

    return max
}

func Max(x, y int) int {
    if x > y {
        return x
    }

    return y
}

func Min(x, y int) int {
    if x > y {
        return y
    }

    return x
}