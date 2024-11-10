给你一个 无重叠的 ，按照区间起始端点排序的区间列表 intervals，其中 intervals[i] = [starti, endi] 表示第 i 个区间的开始和结束，并且 intervals 按照 starti 升序排列。同样给定一个区间 newInterval = [start, end] 表示另一个区间的开始和结束。

在 intervals 中插入区间 newInterval，使得 intervals 依然按照 starti 升序排列，且区间之间不重叠（如果有必要的话，可以合并区间）。

返回插入之后的 intervals。

注意 你不需要原地修改 intervals。你可以创建一个新数组然后返回它。

## 思路
- 不重叠的绿区间，在蓝区间的左边  
- 有重叠的绿区间   
- 不重叠的绿区间，在蓝区间的右边   

![](https://pic.leetcode-cn.com/1604465027-kDWfBc-image.png)

逐个分析
不重叠，需满足：绿区间的右端，位于蓝区间的左端的左边，如 [1,2]。

则当前绿区间，推入 res 数组，指针 +1，考察下一个绿区间。
循环结束时，当前绿区间的屁股，就没落在蓝区间之前，有重叠了，如 [3,5]。
现在看重叠的。我们反过来想，没重叠，就要满足：绿区间的左端，落在蓝区间的屁股的后面，反之就有重叠：绿区间的左端 <= 蓝区间的右端，极端的例子就是 [8,10]。

和蓝有重叠的区间，会合并成一个区间：左端取蓝绿左端的较小者，右端取蓝绿右端的较大者，不断更新给蓝区间。
循环结束时，将蓝区间（它是合并后的新区间）推入 res 数组。
剩下的，都在蓝区间右边，不重叠。不用额外判断，依次推入 res 数组。


```py
class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        n = len(intervals)
        i = 0
        res = []
        while i < n and intervals[i][1] < newInterval[0]:
            res.append(intervals[i])
            i += 1
        while i < n and intervals[i][0] <= newInterval[1]:
            newInterval[0] = min(newInterval[0], intervals[i][0])
            newInterval[1] = max(newInterval[1], intervals[i][1])
            i += 1
        res.append(newInterval)
        while i < n:
            res.append(intervals[i])
            i += 1
        return res
```