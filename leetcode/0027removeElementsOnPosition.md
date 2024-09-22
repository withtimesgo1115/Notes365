给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。

假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：

更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
返回 k。


输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2,_,_]
解释：你的函数函数应该返回 k = 2, 并且 nums 中的前两个元素均为 2。
你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。

输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3,_,_,_]
解释：你的函数应该返回 k = 5，并且 nums 中的前五个元素为 0,0,1,3,4。
注意这五个元素可以任意顺序返回。
你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。


本题考察的是双指针，通过索引来解决。题设非常巧妙，不需要考虑数组后面的部分，只需要保证前面的元素即可。

思路就是分成2部分，第一部分的最后元素的索引+1就是长度。
双指针i, j分别从0和最后一个位置开始移动。相等则直接交换，把要删除的放后面就行了，不等就继续往下走。

```python
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        j = len(nums) - 1
        i = 0
        while i <= j:
            if nums[i] == val:
                self.swap(nums, i, j)
                j -= 1
            else:
                i += 1
        return j + 1
    
    def swap(self, nums, i, j):
        nums[i], nums[j] = nums[j], nums[i]
```