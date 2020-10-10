// 解题思路：DP
// 时间复杂度O(M*N), 空间复杂度O(M*N)
func uniquePaths(m int, n int) int {
    if m < 0 || n < 0 {
        return 0
    }
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    // 初始底部
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    } 

    // dp
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    return  dp[m-1][n-1]
}