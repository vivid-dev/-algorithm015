// 解题思路：DP
// 时间复杂度O(M*N), 空间复杂度O(M*N)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
        return 0
    }
    
    dp := make([][]int, len(obstacleGrid))
    for i := range dp {
        dp[i] = make([]int, len(obstacleGrid[0]))
    }

    for i := 0; i < len(obstacleGrid); i++ {
        if obstacleGrid[i][0] != 0 {
            break
        }
        dp[i][0] = 1
    }

    for i := 0; i < len(obstacleGrid[0]); i++ {
        if obstacleGrid[0][i] != 0 {
            break
        }
        dp[0][i] = 1
    }

    for i := 1; i < len(obstacleGrid); i++ {
        for j := 1; j < len(obstacleGrid[i]); j++ {
            if obstacleGrid[i][j] == 0 {
               dp[i][j] = dp[i-1][j] + dp[i][j-1]             
            }
        }
    }

    return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}