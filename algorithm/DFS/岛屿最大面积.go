/*
给你一个大小为 m x n 的二进制矩阵 grid 。

岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

岛屿的面积是岛上值为 1 的单元格的数目。

计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
*/

func maxAreaOfIsland(grid [][]int) int {
	dirs := [4][2]int{{0,1},{0,-1},{1,0},{-1,0}}
	rows, cols := len(grid), len(grid[0])

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if (x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == 0) {
			return 0
		}
		grid[x][y] = 0
		area := 1
		for _, dir := range dirs {
			area += dfs(x + dir[0], y + dir[1])
		}
		return area
	}
	maxArea := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(i, j))
			}
		}
	}
	return maxArea
}

func max(a, b int) {
	if a > b {
		return a
	}
	return b
}