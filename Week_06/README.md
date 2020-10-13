### 0 前言
动规是一个比较经（zhong）典（yao）的算（tao）法（lu），面试也经常用到，这里做个简单总结。
1 适合的场景
1） 求最值；
2） 计算方案总数；
3） 判断方案是否可行；
### 2 解题思路
首先是看场景是否适合，其次再找递推关系，比如最终是f(x,y) = max(f(x-1, y-1), f(x, y-1), f(x-1, y))，底部状态，比如，x == 0 || y == 0。
注意：
类似于递归，找递推关系时，可以假设第f（x，y）这个事情就做完了，具体要怎么做先不要管，然后根据题意找出上一状态，得到f（x，y）与上一个状态的递推方程。到底部状态时，再得出这个事情怎么做就行。分析由顶到底，底部具体分析要做什么，计算由底部向上计算。
### 3 代码示例
#### 3.1 求最值
https://leetcode.com/problems/longest-palindromic-substring/

```
// 动规适合的场景：
// 1.求最值；2.计算方案总数；3.判断方案是否可行
// 此处符合1，求最值。
// 1 最优值优选动规,用s[l, r]表示以l为起始，r为终止的串
// 2 状态转移方程:
// 分类讨论：
// 1）s[l] == s[r], 
//    若长度为2, 则s[l, r] = 2
//    否则，若s[l+1, r-1]是回文 => s[l, r]是回文串，s[l, r] = s[l+1, r-1] + 2
//    否则，s[l, r] = 0
// 2）s[l] != s[r], s[l, r] = 0
// 3 底部：
//   s[l, r]初始状态为l==r，值为1,即只有一个字符时，自身就是回文串。

class Solution {
public:
    string longestPalindrome(string s) {
        int len=s.size();
        if (len <= 1) {
            return s;
        }
        vector<vector<int> > dp(len,vector<int>(len));
        int max = 1;
        int left = 0;

        // 初始底和整个dp数组，长度为一个字符，初始为1，否则初始为0
        for (int i = 0; i < len; ++i) {
            for (int j = 0; j < len; ++j) {
                if (i == j) {
                    dp[i][j] = 1;
                } else {
                    dp[i][j] = 0;
                }
            }
        }
        // 进行dp计算
        for (int j = 0; j < len; ++j) {
            // 由底向上计算，底部i==j已经被初始化，因此这里用j-1
            for (int i = j - 1; i >= 0; --i) { 
                if (s[i] == s[j]) {
                    if (j - i == 1) {
                        dp[i][j] = 2;
                    } else {
                        if (dp[i + 1][j - 1] != 0) {
                            dp[i][j] = dp[i + 1][j - 1] + 2;
                        } else {
                            dp[i][j] = 0;
                        }
                    }
                } else {
                    dp[i][j] = 0;
                } 
                // 当前值大于最大值，则记录下标
                if (dp[i][j] >= max) {
                    left = i;
                    max = dp[i][j];
                }
                
            }
        }
        
        return s.substr(left, max);
    }
};
```
#### 3.2 计算方案总数
https://leetcode.com/problems/unique-paths/
```
// 动规适合的场景：
// 1.求最值；2.计算方案总数；3.判断方案是否可行
// 此处符合2，计算方案总数。
// 1）状态转移方程
// cnt[i][j]表示到达当前位置的总数，那么不难想到:
// cnt[i][j] = cnt[i-1][j] + cnt[i][j-1]
// 2）底
// i==0 || j == 0，只有1种走法;
class Solution {
public:
    int uniquePaths(int m, int n) {
        if (m == 0 && n == 0) {
            return 0;
        }
        
        vector<vector<int>> cnt(m, vector<int>(n));
        // 初始化底
        for (int i = 0; i < m; ++i) {
            cnt[i][0] = 1;
        }
         for (int j = 0; j < n; ++j) {
             cnt[0][j] = 1;
         }
        // 由底向上进行dp
        for (int i = 1; i < m; ++i) {
            for (int j = 1; j < n; ++j) {
                cnt[i][j] = cnt[i - 1][j] + cnt[i][j - 1];
            }
        }
        return cnt[m - 1][n - 1];
    }
};
```
#### 3.3 判断方案是否可行
略。
