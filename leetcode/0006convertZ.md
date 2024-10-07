将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
 

示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"

## 思路
设 numRows 行字符串分别为 s1, s2, ... , sn 则容易发现：按顺序遍历字符串 s 时，每个字符 c 在 N 字形中对应的 行索引 先从
s1增大至 sn，再从 sn减小至 s1…… 如此反复。

因此解决方案为：模拟这个行索引的变化，在遍历 s 中把每个字符填到正确的行 res[i], 这里用到一个flag来做变向


```py
class Solution:
    def convert(self, s, numRows):
        # 特殊情况
        if numRows < 2: return s
        # 先初始化结果字符串列表，我们就是要来更新每个res[i]
        res = ["" for _ in range(numRows)]
        # 初始化，flag开始为-1，是为了先进行一次换向
        i, flag = 0, -1
        # 遍历一次
        for c in s:
            # 更新res[i]
            res[i] += c
            # 对于边界条件，切换flag的方向
            if i == 0 or i == numRows - 1:
                flag = -flag
            # 每次更新一下i的值
            i += flag
        return "".join(res)
```