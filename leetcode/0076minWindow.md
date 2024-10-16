给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

 

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
 

示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。

## 思路


```py
class Solution:
    def minWindow(self, s: str, t: str) -> str:
        from collections import defaultdict
        # 记录需要的字符和它们的数量
        need = defaultdict(int)
        for char in t:
            need[char] += 1
        
        # 滑动窗口中的字符记录
        window = defaultdict(int)
        left, right = 0, 0
        count = 0  # 记录窗口中满足要求的字符个数
        start = 0  # 记录最小窗口的起始位置
        min_len = float('inf')  # 用来记录最小窗口的长度

        while right < len(s):
            # 更新窗口
            c = s[right]
            right += 1
            if c in need:
                window[c] += 1
                if window[c] == need[c]:
                    count += 1
            
            # 收缩左边界，找到更小的符合条件的子串
            while count == len(need):
                # 更新最小长度和起始位置
                if right - left < min_len:
                    min_len = right - left
                    start = left
                
                d = s[left]
                left += 1
                if d in need:
                    if window[d] == need[d]:
                        count -= 1
                    window[d] -= 1
        return "" if min_len == float('inf') else s[start:start + min_len]
```