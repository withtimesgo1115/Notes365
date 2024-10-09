给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：

如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？


示例 1：

输入：s = "abc", t = "ahbgdc"
输出：true
示例 2：

输入：s = "axc", t = "ahbgdc"
输出：false

## 思路
使用双指针，分别从第一个元素开始遍历，贪心查找。  
只有当两个元素相等时，i, j均移动到下一个元素，继续查找  
否则，只有j+=1，i不变，继续查找


```py
class Solution:
    def isSubsequence(self, s, t):
        n, m = len(s), len(t)
        i = j = 0
        while i < n and j < m:
            if s[i] == t[j]:
                i += 1
            j += 1
        # 返回i == n 很巧妙，如果能找到子序列，那么
        return i == n
```