给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。

以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。


 
示例 1：

输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
示例 2：

输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。
示例 3：

输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

## 思路
双指针方法，为什么双指针往中间移动时，不会漏掉某些情况呢？要解答这个问题，我们要从 缩减搜索空间 的角度思考这个解法。  
需要注意的是，虽然本题叫做 Two Sum II，但解法和 Two Sum 完全不同。  

在这道题中，我们要寻找的是符合条件的一对下标 (i,j)，它们需要满足的约束条件是：
- i、j 都是合法的下标，即 0 ≤ i < n, 0 ≤ j < n
- i < j（题目要求）
![](https://pic.leetcode-cn.com/6ee3750f6036a7a6249197e5b640bfc0564153ca1a61c1e35aad51f3a8f9dc5e.jpg)

由于 i、j 的约束条件的限制，搜索空间是白色的倒三角部分。可以看到，搜索空间的大小是 O(n2) 数量级的。如果用暴力解法求解，一次只检查一个单元格，那么时间复杂度一定是 O(n2)。要想得到 O(n) 的解法，我们就需要能够一次排除多个单元格。那么我们来看看，本题的双指针解法是如何削减搜索空间的：

假设此时 A[0] + A[7] 小于 target。这时候，我们应该去找和更大的两个数。由于 A[7] 已经是最大的数了，其他的数跟 A[0] 相加，和只会更小。也就是说 A[0] + A[6] 、A[0] + A[5]、……、A[0] + A[1] 也都小于 target，这些都是不合要求的解，可以一次排除。这相当于 i=0 的情况全部被排除。对应用双指针解法的代码，就是 i++，对应于搜索空间，就是削减了一行的搜索空间，如下图所示。

![](https://pic.leetcode-cn.com/50d93bb2d2ce3e2985460586d4350e8205543965d9689632a20f5650dde3cb95.jpg)

排除掉了搜索空间中的一行之后，我们再看剩余的搜索空间，仍然是倒三角形状。我们检查右上方的单元格 (1,7)，计算 A[1] + A[7] 与 target 进行比较。

![](https://pic.leetcode-cn.com/3e305bd710d6f2c3730bd3050f49439f9e63b19eee24066f6642c393df6fdafb.jpg)

假设此时 A[0] + A[7] 大于 target。这时候，我们应该去找 和更小的两个数。由于 A[1] 已经是当前搜索空间最小的数了，其他的数跟 A[7] 相加的话，和只会更大。也就是说 A[1] + A[7] 、A[2] + A[7]、……、A[6] + A[7] 也都大于 target，这些都是不合要求的解，可以一次排除。这相当于 j=n 的情况全部被排除。对应用双指针解法的代码，就是 j--，对应于搜索空间，就是削减了一列的搜索空间


可以看到，无论 A[i] + A[j] 的结果是大了还是小了，我们都可以排除掉一行或者一列的搜索空间。经过 n 步以后，就能排除所有的搜索空间，检查完所有的可能性。搜索空间的减小过程如下面动图所示：
![](https://pic.leetcode-cn.com/9ebb3ff74f0706c3c350b7fb91fea343e54750eb5b6ae6a4a3493421a019922a.gif)



```py
class Solution:
    def twoSum(self, numbers, target):
        i = 0
        j = len(numbers) - 1
        while i < j:
            cur_sum = numbers[i] + numbers[j]
            if cur_sum < target:
                i += 1
            elif cur_sum > target:
                j -= 1
            else:
                return [i+1, j+1]
        return [-1, -1]
```