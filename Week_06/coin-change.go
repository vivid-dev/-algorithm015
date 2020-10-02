// 解题思路：DP
// 时间复杂度O(M*N), 空间复杂度O(N)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount+1
	}
    dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := range coins {
			if i-coins[j] < 0 || dp[i-coins[j]] == amount + 1 {
				continue
			}
			dp[i] = min(dp[i], 1 + dp[i-coins[j]])
		}
	}

	if dp[amount] == amount + 1 {
		return -1
	}

	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}