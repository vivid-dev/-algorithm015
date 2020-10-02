// 解题思路：动态规划
// 时间复杂度:O(N^2), 空间复杂度O(N)
func jump(nums []int) int {
    // init dp结构
    n := len(nums) - 1
    dp := make([]int, n + 1)
    
    // 初始化底
    for i := n - 1; i >= 0; i-- {
        if nums[i] >= n - i {
            dp[i] = 1
        }
    }

    // dp
    for i := n - 1; i >= 0; i-- {
        for j := i - 1; j >= 0; j-- {
			// 从终点跳过来的 && 可以跳过去  && (没走过 || 有更优值)
            if dp[i] != 0 && nums[j] >= i - j && (dp[j] == 0 || dp[j] >= dp[i] + 1) {
                dp[j] = dp[i] + 1
            }
        }
    }

    return dp[0]
}