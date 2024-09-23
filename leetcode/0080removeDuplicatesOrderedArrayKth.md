给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

 

说明：

为什么返回数值是整数，但输出的答案是数组呢？

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
 

示例 1：

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,1,2,3,3]
输出：7, nums = [0,0,1,1,2,3,3]
解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为 0, 0, 1, 1, 2, 3, 3。不需要考虑数组中超出新长度后面的元素。

思路：通过移动快指针 fast 来找到新的元素，然后将新元素复制到慢指针 slow 的位置，从而在原地删除重复项。当找到一个与 nums[slow - k] 不同的元素时，我们就将新元素复制到 nums[slow]. `

```py
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        k = 2  # 每个元素最多出现的次数
        if nums is None or len(nums) <= k: return len(nums)

        slow, fast = k, k
        while fast < len(nums):
            # 如果 nums[fast] 不等于 nums[slow - k]
            # 则将 nums[fast] 复制到 nums[slow]，并将 slow 向前移动一位
            # nums[slow - k] 是当前考虑的元素在新数组中的第一个可能的位置
            if nums[fast] != nums[slow - k]:
                nums[slow] = nums[fast]
                slow += 1
            # 移动fast指针，检查下一个元素
            fast += 1
        # 因为slow已经向前移动了一位，所以返回索引就是长度
        return slow
```