package main

func main() {
	// grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	// grid = [][]int{{0, 1, 1, 0, 0, 0}, {0, 1, 0, 1, 1, 0}, {0, 1, 1, 0, 1, 0}, {0, 0, 0, 1, 1, 0}, {1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 0}}
	// grid := [][]int{{0, 0, 0, 0}, {1, 0, 0, 1}, {0, 1, 0, 0}, {0, 0, 0, 0}}
}

var steps = [][]int{{1, 1}, {1, 0}, {0, 1}, {1, -1}, {-1, 1}, {-1, -1}, {-1, 0}, {0, -1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}
	queue := [][]int{{0, 0}}
	grid[0][0] = 1
	for len(queue) > 0 {
		level := len(queue)
		for node := 0; node < level; node++ {
			i, j := queue[node][0], queue[node][1]
			for _, step := range steps {
				newi, newj := i+step[0], j+step[1]
				if 0 <= newi && newi < n && 0 <= newj && newj < n {
					if grid[newi][newj] == 0 {
						queue = append(queue, []int{newi, newj})
						grid[newi][newj] = grid[i][j] + 1
					}
					if newi == n-1 && newj == n-1 {
						return grid[newi][newj]
					}
				}
			}
		}
		queue = queue[level:]
	}
	return -1
}

func shortestPathBinaryMatrixError(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	INF := n*m + 1
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == 1 {
			return INF
		}
		if grid[i][j] < 0 {
			return -grid[i][j]
		}
		defer func() {
			// fmt.Println(i, j, grid[i][j], res)
			grid[i][j] = -res
		}()
		if i == n-1 && j == m-1 {
			return 1
		}
		res = INF
		grid[i][j] = 1
		for _, step := range steps {
			x, y := i+step[0], j+step[1]
			res = min(res, dfs(x, y)+1)
		}
		grid[i][j] = 0
		return res
	}
	ans := dfs(0, 0)
	if ans == INF {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shortestPathBinaryMatrix20230526(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[n-1][m-1] == 1 {
		return -1
	}
	valid := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m && grid[i][j] == 0
	}
	q := make([][2]int, 0)
	q = append(q, [2]int{0, 0})
	grid[0][0] = 1
	level := 1
	for len(q) > 0 {
		s := len(q)
		for i := 0; i < s; i++ {
			x, y := q[0][0], q[0][1]
			if x == n-1 && y == m-1 {
				return level
			}
			for _, step := range steps {
				xi, yi := x+step[0], y+step[1]
				if valid(xi, yi) {
					grid[xi][yi] = 1
					q = append(q, [2]int{xi, yi})
				}
			}
			q = q[1:]
		}
		level++
	}
	return -1
}
