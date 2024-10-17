根据 百度百科 ， 生命游戏 ，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。

![](https://assets.leetcode.com/uploads/2020/12/26/grid1.jpg)

![](https://assets.leetcode.com/uploads/2020/12/26/grid2.jpg)

## 思路


```py
class Solution:
    def gameOfLife(self, board):
        import numpy as np
        r, c = len(board), len(board[0])
        # 下面两行做 zero padding
        board_exp = np.array([[0 for _ in range(c+2)]for _ in range(r+2)])
        board_exp[1:1+r, 1:1+c] = np.array(board)
        # 设置卷积核
        kernel = np.array([[1,1,1],[1,0,1],[1,1,1]])
        # 开始卷积
        for i in range(1, r+1):
            for j in range(1, c+1):
                # 统计细胞周围 8 个位置的状态
                temp_sum = np.sum(kernel * board_exp[i-1:i+2, j-1:j+2])
                # 按照题目规则进行判断
                if board_exp[i, j] == 1:
                    if temp_sum < 2 or temp_sum > 3:
                        board[i-1][j-1] = 0
                else:
                    if temp_sum == 3:
                        board[i-1][j-1] = 1
```