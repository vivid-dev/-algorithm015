// dp
// 解题思路：
// 底：dp[0] = 0，dp[1] = 1
// dp[n] = min(dp[a]*2, dp[b]*3, dp[c]*5)
// 复杂度: 时间复杂度O(N), 空间复杂度O(N)
func nthUglyNumber(n int) int {
    dp := make([]int, n + 1)
    dp[0] = 0
    dp[1] = 1
    // 有底向上dp
    for a, b, c, i:= 1, 1, 1, 2; i <= n; i++ {
        dp[i] = min(dp[a]*2, min(dp[b]*3, dp[c]*5))
        if dp[i] == dp[a]*2 {
            a++
        }
        if dp[i] == dp[b]*3 {
            b++
        }
        if dp[i] == dp[c]*5 {
            c++
        }
    }
    
    return dp[n]
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}