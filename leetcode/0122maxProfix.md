给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。

 

示例 1：

输入：prices = [7,1,5,3,6,4]
输出：7
解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
最大总利润为 4 + 3 = 7 。

## 思路
1. 定义状态
状态 dp[i][j] 定义如下：

dp[i][j] 表示到下标为 i 的这一天，持股状态为 j 时，我们手上拥有的最大现金数。

注意：限定持股状态为 j 是为了方便推导状态转移方程，这样的做法满足 无后效性。

其中：

第一维 i 表示下标为 i 的那一天（ 具有前缀性质，即考虑了之前天数的交易 ）；   
第二维 j 表示下标为 i 的那一天是持有股票，还是持有现金。   
这里 0 表示持有现金（cash），1 表示持有股票（stock）。   

2. 思考状态转移方程

状态从持有现金（cash）开始，到最后一天我们关心的状态依然是持有现金（cash）；
每一天状态可以转移，也可以不动。状态转移用下图表示：


（状态转移方程写在代码中）

说明：

由于不限制交易次数，除了最后一天，每一天的状态可能不变化，也可能转移；
写代码的时候，可以不用对最后一天单独处理，输出最后一天，状态为 0 的时候的值即可。  

3. 确定初始值
起始的时候：

如果什么都不做，dp[0][0] = 0；
如果持有股票，当前拥有的现金数是当天股价的相反数，即 dp[0][1] = -prices[i]；


4. 确定输出值
终止的时候，上面也分析了，输出 dp[len - 1][0]，因为一定有 dp[len - 1][0] > dp[len - 1][1]。



```py
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        n = len(prices)
        if n < 2:
            return 0
        
        dp = [[0] * 2 for _ in range(n)]
        dp[0][0] = 0
        dp[0][1] = -prices[0]

        for i in range(1, n):
            # 持有股票时，可以卖出赚钱
            dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
            # 持有现金时，可以买入股票
            dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
        # 返回最后一天，手里有现金时的总数
        return dp[n-1][0]
```