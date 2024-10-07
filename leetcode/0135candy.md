n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。

 

示例 1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
示例 2：

输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。


## 思路
规则定义： 设学生 A 和学生 B 左右相邻，A 在 B 左边；
左规则： 当 ratingsB > ratingsA时，B 的糖比 A 的糖数量多。
右规则： 当 ratingsA > ratingsB时，A 的糖比 B 的糖数量多。   

相邻的学生中，评分高的学生必须获得更多的糖果 等价于 所有学生满足左规则且满足右规则。


1. 先从左到右遍历ratings，并记录到left中
    - 先给每个学生1个糖
    - 如果ratingsi > ratingsi-1 则i比i-1多1个
    - 如果ratingsi <= ratingsi-1 则不做处理，留给第二次遍历处理
2. 然后遍历ratings, 从后往前，记录到right
3. 取2次遍历对应糖果的最大值

```py
class Solution:
    def candy(self, ratings):
        left = [1 for _ in range(len(ratings))]
        right = left[:]
        for i in range(1, len(ratings)):
            if rating[i] > ranting[i-1]:
                left[i] = left[i-1] + 1
        count = left[-1]
        for i in range(len(ratings) - 1 - 1, -1, -1):
            if ratings[i] > ratings[i+1]:
                right[i] = right[i+1] + 1
            count += max(left[i], right[i])
        return count
```