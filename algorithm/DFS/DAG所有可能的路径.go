/*
给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）

graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。

输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
*/

func allPaths(graph [][]int) [][]int {
	var res [][]int // 二维数组存放结果
	var path []int  // 存放当前路径
	var dfs func(int) // dfs函数

	dfs = func(x int) {
		// 终止条件，就是找到了最后一个位置
		if x == len(graph) - 1 {
			temp := make([]int, len(path))
			copy(temp, path) // 注意要复制一下，不能直接用path，因为后面可能会修改
			res = append(res, temp)
			return
		}
		for _, next := range graph[x] { // for循环检查当前节点的邻接元素
			path = append(path, next)   // 把邻接的元素加到Path里
			dfs(next)					// 递归
			path = path[:len(path)-1]  // 回溯到上一步，继续看其他邻接元素
		}
	}

	path = append(path, 0) // 起始节点为0
	dfs(0)  // 从0开始dfs
	return res
}
