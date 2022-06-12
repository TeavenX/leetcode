package main

import (
	"fmt"
	"sort"
)

func main() {
	forest := [][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}}
	fmt.Println(cutOffTree(forest))
}

var walks = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func cutOffTree(forest [][]int) int {
	n, m := len(forest), len(forest[0])
	trees := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if forest[i][j] > 1 {
				// 将每一颗树都压到切片中
				trees = append(trees, []int{forest[i][j], i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool {
		// 这里按照树的高度进行排序
		return trees[i][0] < trees[j][0]
	})
	var bfs func(startX, startY, targetX, targetY int) int
	bfs = func(startX, startY, targetX, targetY int) int {
		// 广度优先搜索：
		// 		在已知当前坐标的情况下，寻找到达目标坐标所需要的步数
		step := 0
		visited := make(map[int]bool)
		// 因为树的先后顺序与位置不确定，可能需要来回走，所以visited在每次寻找的时候初始化（也就是当前寻找任务下，不重复走）
		// 切片不能作为map的key，刚好题目说明了树高不会重复，因此可以用树高作为是否访问过的判断依据
		// 因为存在来回走的情况，所以类似「岛屿数量」题目在原地修改的方式不可用，这里的空间开销无法优化
		visited[forest[startX][startY]] = true // 初始化起始坐标为已访问
		queue := [][]int{{startX, startY}}     // 将初始坐标放进队列
		for len(queue) > 0 {
			size := len(queue)
			for i := 0; i < size; i++ {
				x, y := queue[i][0], queue[i][1]
				if x == targetX && y == targetY {
					// 找到了目标位置，返回结果
					return step
				}
				for _, walk := range walks {
					newX, newY := x+walk[0], y+walk[1]
					if newX >= 0 && newX < n && newY >= 0 && newY < m && forest[newX][newY] != 0 && !visited[forest[newX][newY]] {
						// 上下左右坐标在数组范围内
						// 并且新坐标对应的树高不是0（不是障碍）
						// 并且新坐标没有访问过
						queue = append(queue, []int{newX, newY})
						visited[forest[newX][newY]] = true
					}
				}
			}
			step++
			queue = queue[size:]
		}
		// 执行到这个return意味着上面搜索失败，无法到达目标坐标
		return -1
	}
	result := 0
	preX, preY := 0, 0
	for _, tree := range trees {
		// 前面已经按照树高进行排序，所以这里直接遍历
		step := bfs(preX, preY, tree[1], tree[2])
		if step == -1 {
			// step=-1意味着有一颗树无法到达，结果返回-1
			return -1
		}
		result += step
		preX, preY = tree[1], tree[2]
	}
	return result
}
