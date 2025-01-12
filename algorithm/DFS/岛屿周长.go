/*
给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。

网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。
*/

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	rows := len(grid)
	cols := len(grid[0])
	if rows == 0 || cols == 0 {
		return 0
	}
	var dfs func ([][]int, int, int) int

	dfs = func(grid [][]int, x int, y int) int {
		if (x < 0 || x >= rows || y < 0 || y >= cols) {
			return 1
		}
		if grid[x][y] == 0 {
			return 1
		}
		if grid[x][y] == 2 {
			return 0
		}
		grid[x][y] = 2
		return dfs(grid, x - 1, y) + dfs(grid, x + 1, y) + dfs(grid, x, y - 1) + dfs(grid, x, y + 1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				perimeter = dfs(grid, i, j)
			}
		}
	}
	return perimeter
}