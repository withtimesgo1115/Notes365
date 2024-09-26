给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

 

示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。


## 思路
![](https://pic.leetcode-cn.com/2a834dafa7bf590df1413fc742b07099854b6c6b842a5f7677564ccd044b5d69.png)

字符串前后有""符号，所以要注意开头和结尾的索引

1. 初始化 dp=[False,⋯,False]，长度为 n+1。n 为字符串长度。dp[i] 表示 s 的前 i 位是否可以用 wordDict 中的单词表示。
2. 初始化 dp[0]=True，空字符可以被表示。
3. 遍历字符串的所有子串，遍历开始索引 i，遍历区间[0,n)：
遍历结束索引 j，遍历区间 [i+1,n+1)：
    若 dp[i]=True 且 s[i,⋯,j) 在 wordlist 中：dp[j]=True。解释：dp[i]=True 说明 s 的前 i 位可以用 wordDict 表示，则 s[i,⋯,j) 出现在 wordDict 中，说明 s 的前 j 位可以表示。
4. 返回 dp[n]

```py
class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        n = len(s)
        dp = [False] * (n + 1)
        dp[0] = True
        for i in range(n):
            for j in range(i+1, n+1):
                if dp[i] and s[i:j] in wordDict:
                    dp[j] = True
        return dp[-1]
```