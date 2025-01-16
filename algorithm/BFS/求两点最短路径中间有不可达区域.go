/* 
有n✖️m的网格, 现在你在（1,1），要走到(n,m) 
在(x,y)处有个激光, 与(x,y)距离在d以内包括d的位置都会受伤。要避开这个激光，问最短路径是多少？
如果走不到的话输出-1，每次只能往上下左右走一步 距离计算 |x-a| + |y-b|
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func isValid(x, y, n, m int) bool {
	return x >= 1 && x <= n && y >= 1 && y <= m 
}

func isRaserArea(x, y, laserX, laserY, d int) bool {
	return math.Abs(float64(x-laserX)) + math.Abs(float64(y-laserY)) <= float64(d)
}

func shortestPath(n, m, laserX, laserY, d int) int {
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	// 创建一个队列用于BFS 
	visited := make([][]bool, n+1)
	for i := range visited {
		visited[i] = make([]bool, m+1)
	}
	
	// 创建队列用于BFS
	// 把起始位置加入队列，同时染色，标志已经访问过
	queue := []Point{{1,1}}
	visited[1][1] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			// 当前层的节点个数
			cur := queue[0]
			queue = queue[1:]
			if cur.x == n && cur.y == m {
				return steps
			}

			for _, dir := range directions {
				newX, newY := cur.x + dir.x, cur.y + dir.y
				newPoint := Point{newX, newY}
				if isValid(newX, newY, n, m) && !isRaserArea(newX, newY, laserX, laserY, d) && !visited[newX][newY] {
					queue = append(queue, newPoint)
					visited[newX][newY] = true
				}
			}
		}
		steps++
	}

	return -1
}

func main() {
	n, m := 5, 5
	x, y := 3, 3
	d := 1
	result := shortestPath(n, m, x, y, d)
	fmt.Println(result)
}