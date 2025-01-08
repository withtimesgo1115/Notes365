/*
给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）

graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。

输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
*/

func allPaths(graph [][]int) [][]int {
	var res [][]int
	var path []int
	var dfs func(int)

	dfs = func(x int) {
		if x == len(graph) - 1 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, next := range graph[x] {
			path = append(path, next)
			dfs(next)
			path = path[:len(path)-1]
		}
	}

	path = append(path, 0)
	dfs(0)
	return res
}
