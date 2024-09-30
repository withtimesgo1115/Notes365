给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

## 思路
动态规划一般分为一维、二维、多维（使用 状态压缩），对应形式为 dp(i)、dp(i)(j)、二进制dp(i)(j)。

1. 动态规划做题步骤

明确 dp(i) 应该表示什么（二维情况：dp(i)(j)）；  
根据 dp(i) 和 dp(i−1) 的关系得出状态转移方程；   
确定初始条件，如 dp(0)。   

2. 本题思路

其实方法一的思路不是凭空想象的，而是由动态规划的思想演变而来。这里介绍一维动态规划思想。

dp[i] 表示前 i 天的最大利润，因为我们始终要使利润最大化，则：

dp[i]=max(dp[i−1],prices[i]−minprice)


 ```py
class Solution:
    def maxProfit(self, prices):
        n = len(prices)
        if n == 0:
            return 0
        dp = [0] * n
        minprice = prices[0]

        for i in range(1, n):
            minprice = min(minprice, prices[i])
            dp[i] = max(dp[i-1], prices[i] - minprice)
        return dp[-1]
 ```