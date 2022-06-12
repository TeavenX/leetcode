package main

func main() {

}

type node struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	dirs := []node{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	queue := make([]node, 0)
	result := 0
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, node{i, j})
				visited[i][j] = true
			}
		}
	}
	for len(queue) > 0 {
		lenQ := len(queue)
		cacheResult := false
		for i := 0; i < lenQ; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				x, y := cur.x+dir.x, cur.y+dir.y
				if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && grid[x][y] == 1 {
					grid[x][y] = 2
					queue = append(queue, node{x, y})
					visited[x][y] = true
					cacheResult = true
				}
			}
		}
		if cacheResult {
			result++
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				result = -1
				break
			}
		}
	}
	return result
}
