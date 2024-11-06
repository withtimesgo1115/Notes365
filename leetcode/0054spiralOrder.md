给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

![](https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg)

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

这里的方法不需要记录已经走过的路径，所以执行用时和内存消耗都相对较小

1. 首先设定上下左右边界
2. 其次向右移动到最右，此时第一行因为已经使用过了，可以将其从图中删去，体现在代码中就是重新定义上边界
3. 判断若重新定义后，上下边界交错，表明螺旋矩阵遍历结束，跳出循环，返回答案
4. 若上下边界不交错，则遍历还未结束，接着向下向左向上移动，操作过程与第一，二步同理
5. 不断循环以上步骤，直到某两条边界交错，跳出循环，返回答案


```py
class Solution:
    def spiralOrder(self, matrix):
        ans = []
        if not matrix:
            return ans
        # 边界索引
        u, b = 0, len(matrix) - 1
        l, r = 0, len(matrix[0]) - 1

        while True:
            # 向右移动
            for i in range(l, r + 1):
                ans.append(matrix[u][i])
            # 重新设置上边界
            u += 1
            if u > b:
                break
            # 向下移动
            for i in range(u, b + 1):
                ans.append(matrix[i][r])
            # 重新设置右边界
            r -= 1
            if r < l:
                break
            # 向左移动
            for i in range(r, l - 1, -1):
                ans.append(matrix[b][i])
            # 重新设置下边界
            b -= 1
            if b < u:
                break
            # 向上移动
            for i in range(b, u - 1, -1):
                ans.append(matrix[i][l])
            # 重新设置左边界
            l += 1
            if l > r:
                break
        return ans
```