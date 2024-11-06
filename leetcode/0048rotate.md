给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

![](https://assets.leetcode.com/uploads/2020/08/28/mat1.jpg)

先上下翻转，再左斜对角线翻转感觉更简单些。

```py
class Solution:
    def rotate(self, matrix):
        n = len(matrix)
        # 上下翻转
        for i in range(n // 2):
            for j in range(n):
                matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
        # 左斜对角(\)翻转
        for i in range(n):
            for j in range(i):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
```