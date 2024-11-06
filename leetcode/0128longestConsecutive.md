给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。


示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9


```py
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        res = 0
        hash_map = {}
        for num in nums:
            # 新进来哈希表一个数
            if num not in hash_map:
                # 获取当前数的最左边连续长度,没有的话就更新为0
                left = hash_map.get(num-1, 0)
                # 同理获取右边的数
                right = hash_map.get(num+1, 0)
                """不用担心左边和右边没有的情况
                因为没有的话就是left或者right0
                并不改变什么
                """
                hash_map[num] = 1 # 把当前数加入哈希表，代表当前数字出现过
                # 更新长度
                length = left + 1 + right
                res = max(res, length)
                # 更新最左端点的值，如果left=n存在，那么证明当前数的前n个都存在哈希表中
                hash_map[num-left] = length
                # 更新最右端点的值，如果right=n存在，那么证明当前数的后n个都存在哈希表中
                hash_map[num+right] = length
                # 此时 【num-left，num-right】范围的值都连续存在哈希表中了
                # 只需要端点存值就行，因为端点中的值在遍历的时候如果在哈希表中就会略过
                
                # 即使left或者right=0都不影响结果
        return res
```