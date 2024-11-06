以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

 

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。


```py
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        res = []
        i = 0 # 第一个指针
        intervals.sort() # 先排序，方便处理

        while i < len(intervals):
            t = intervals[i][1] # 当前较大值
            j = i + 1  # 双指针
            # 遍历后续元素，找需要合并的情况
            while j < len(intervals) and intervals[j][0] <= t:
                # 更新当前最大值
                t = max(t, intervals[j][1])
                j += 1
            # 插入需要合并的最大范围
            res.append([intervals[i][0], t])
            i = j # 更新第一个指针
        return res
```