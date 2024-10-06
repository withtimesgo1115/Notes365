给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

 

示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]
 

 ## 思路
初始化结果数组 answer：

1. 首先创建一个与 nums 大小相同的数组 answer，并初始化为全1，因为在计算乘积的过程中初始乘积值是1。

2. 定义左右指针 left 和 right，及左右乘积变量 lp 和 rp：

left 和 right 分别是遍历数组的左、右指针，初始时，left 从0开始，right 从数组的末尾开始。
lp 和 rp 是用来分别保存从左往右和从右往左的累积乘积，初始值为1，因为乘积从1开始。

3. 遍历数组，计算左右乘积：

通过一个 while 循环，left 从左向右遍历，right 从右向左遍历。
对于每个位置的元素 answer[left] 和 answer[right]，分别乘以当前的左乘积 lp 和右乘积 rp，这相当于把当前位置的左右乘积计算到 answer 中。
然后更新左乘积 lp 和右乘积 rp，分别乘以当前的 nums[left] 和 nums[right] 元素。
最后移动左右指针：left 右移，right 左移，直到遍历结束。

4. 返回结果：

最终返回 answer，其中每个位置的值就是除了当前位置以外其他所有元素的乘积。

这个解法之所以不会把自身也乘进去，是因为在计算的过程中，左乘积和右乘积在不同的时间点分别计算，每一步乘积都跳过了当前的元素。具体解释如下：

详细原理
左乘积 lp 和右乘积 rp 的作用：

左乘积 lp 保存的是当前元素左边所有元素的乘积。
右乘积 rp 保存的是当前元素右边所有元素的乘积。
在每一步的计算中，answer[i] 先乘以 lp，然后乘以 rp，这样 answer[i] 就等于除了当前元素 nums[i] 之外其他所有元素的乘积。

```py
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        # 初始化，都为1方便乘积
        answer = [1] * len(nums)
        # 遍历用的双指针
        left, right = 0, len(nums) - 1
        # 累乘变量
        lp, rp = 1, 1
        # 一次遍历，从两边开始
        while right >= 0 and left < len(nums):
            # 把当前位置的左右乘积计算到 answer 中
            answer[right] *= rp
            answer[left] *= lp
            # 乘当前元素
            lp *= nums[left]
            rp *= nums[right]
            left += 1
            right -= 1
        return answer
```