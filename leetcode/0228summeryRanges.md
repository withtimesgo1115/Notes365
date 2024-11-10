给定一个  无重复元素 的 有序 整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b


## 思路
- 双指针

我们可以用双指针 i 和 j 找出每个区间的左右端点。

遍历数组，当 j+1<n 且 nums[j+1]=nums[j]+1 时，指针 j 向右移动，否则区间 [i,j] 已经找到，将其加入答案，然后将指针 i 移动到 j+1 的位置，继续寻找下一个区间。


```py
class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        def f(i, j):
            return str(i) if i == j else f"{nums[i]}->{nums[j]}"
        i = 0
        n = len(nums)
        res = []
        while i < n:
            j = i
            while j < n - 1 and nums[j+1] = nums[j] + 1:
                j += 1
            res.append(f(i, j))
            i = j + 1
        return res
```