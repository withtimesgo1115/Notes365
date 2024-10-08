给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

![](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg)

输入：[1,8,6,2,5,4,8,3,7]
输出：49 
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

## 思路
移动i, j本质上就是排除了柱子，为什么能够排除？  
原因在于高度由低的柱子决定，如果移动高的柱子，宽度变小了，且高度不会高于低的柱子，所以总面积一定减小，所以不要移动较高的柱子，而是移动较低的柱子，来尝试后面的值。

```py
class Solution:
    def maxArea(self, height):
        i = 0
        j = len(height) - 1
        res = 0
        while i < j:
            area = (j - i) * min(height[i], height[j])
            res = max(res, area)
            if height[i] < height[j]:
                i += 1
            else:
                j -= 1
        return res
```