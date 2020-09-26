func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    // init dp
    // dp[i][0] 第i天持有股票的最大收入
    // dp[i][1] 第i天套现股票的最大收入
    dp := make([][2]int, len(prices))
    
    // 初始化底部
    dp[0][0] = -prices[0]
    dp[0][1] = 0

    // dp
    for i := 1; i < len(prices); i++ {
        // 第i天持股收入 = max(第i-1天持股收入，第i-1天套现 & 第i天买入)
        dp[i][0] = max(dp[i-1][0], dp[i-1][1] - prices[i])
        // 第i天套现收入 = max(第i-1天套现收入，第i-1天持股 & 第i天套现)
        dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i])
    }

    // 返回对应天套现的收入
    return dp[len(prices)-1][1]
}

func max(x, y int) int {
    if x > y {
        return x
    }

    return y
}