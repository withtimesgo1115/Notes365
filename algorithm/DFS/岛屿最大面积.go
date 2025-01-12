/*
����һ����СΪ m x n �Ķ����ƾ��� grid ��

���� ����һЩ���ڵ� 1 (��������) ���ɵ���ϣ�����ġ����ڡ�Ҫ������ 1 ������ ˮƽ������ֱ���ĸ������� ���ڡ�����Լ��� grid ���ĸ���Ե���� 0������ˮ����Χ�š�

���������ǵ���ֵΪ 1 �ĵ�Ԫ�����Ŀ��

���㲢���� grid �����ĵ�����������û�е��죬�򷵻����Ϊ 0 ��
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