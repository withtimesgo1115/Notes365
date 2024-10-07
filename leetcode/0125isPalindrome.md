如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

字母和数字都属于字母数字字符。

给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

## 思路
1. 先去掉首尾空格
2. 特殊条件判断
3. 对字符串格式化处理
4. 使用双指针，从头部和尾部分别遍历
5. while循环，i<=j
6. 不符合时直接返回
7. 默认返回True

```py
class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = s.strip()
        if s == "": return True
        s = ''.join(char for char in s.lower() if char.isalnum())
        i, j = 0, len(s) - 1
        while i <= j:
            if s[i] != s[j]:
                return False
            i += 1
            j -= 1
        return True
```