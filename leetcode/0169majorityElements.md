给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

 
示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2

## 思路
其实可以用map的方法来记录，但是有更好的摩尔投票法

核心思想--抵消原则： 在一个数组中，如果某个元素的出现次数超过了数组长度的一半，那么这个元素与其他所有元素一一配对，最后仍然会剩下至少一个该元素。 通过“投票”和“抵消”的过程，可以逐步消除不同的元素，最终留下的候选人就是可能的主要元素。


```py
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        votes, count = 0, 0
        for num in nums:
            if votes == 0:
                x = num
            votes += 1 if num == x else -1
        
        for num in nums:
            if num == x:
                count += 1
        return x if count > len(nums) // 2 else 0
```
