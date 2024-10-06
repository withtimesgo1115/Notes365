编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

 

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

## 思路
1. 当字符串数组长度为 0 时则公共前缀为空，直接返回；
2. 令最长公共前缀 ans 的值为第一个字符串，进行初始化；
3. 遍历后面的字符串，依次将其与 ans 进行比较，两两找出公共前缀，最终结果即为最长公共前缀；
4. 如果查找过程中出现了 ans 为空的情况，则公共前缀不存在直接返回；


```py
class Solution:
    def longestCommonPrefix(self, strs):
        if len(strs) == 0: return ""
        # 以第一个为baseline
        ans = strs[0]
        # 从第二个开始遍历
        for i in range(1, len(strs)):
            # 二层遍历字符的索引
            j = 0
            # 同时检查ans和strs[i][j]
            while j < len(ans) and j < len(strs[i]):
                # 不等说明需要跳出循环了
                if ans[j] != strs[i][j]:
                    break
                # 循环
                j += 1
            # 更新ans
            ans = ans[:j]
            # 如果为空，说明没有公共前缀
            if len(ans) == 0:
                return ""

        return ans
```