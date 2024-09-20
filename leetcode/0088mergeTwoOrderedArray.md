给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 


### 这个题目考察的是三指针、就地修改、有序数组、倒序遍历  

思路的重点一个是从后往前确定两组中该用哪个数字  
另一个是结束条件以第二个数组全都插入进去为止，因为正常遍历一遍后，nums1中大的值都已经确认完毕，还需要再看下nums2是否还有剩余没有加进去的值。

1. 倒序遍历，初始化两个指针
2. 第三个指针用于更新数组
3. 结束条件以第二个数组全都插入进去为止
4. while循环

```python
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        i = m - 1
        j = n - 1
        k = m + n - 1

        while i >= 0 and j >= 0:
            if nums1[i] > nums2[j]:
                nums1[k] = nums1[i]
                i -= 1
            else:
                nums1[k] = nums2[j]
                j -= 1
            k -= 1
        while j >= 0:
            num1[k] = nums2[j]
            j -= 1
            k -= 1
```