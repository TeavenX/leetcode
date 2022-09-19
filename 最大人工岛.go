package main

func main() {

}

var steps = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func largestIsland(grid [][]int) int {
	n := len(grid)
	cache := make(map[int]int)
	var dfs func(i, j, idx int) int
	dfs = func(i, j, idx int) int {
		if i < 0 || j < 0 || i >= n || j >= n {
			return 0
		}
		area := 0
		if grid[i][j] == 1 {
			grid[i][j] = idx
			area++
			for _, step := range steps {
				area += dfs(i+step[0], j+step[1], idx)
			}
		}
		return area
	}
	size := 0
	idx := 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			area := dfs(i, j, idx)
			if area > 0 {
				cache[idx] = area
				size = max(size, area)
				idx++
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				area := 1
				visited := make(map[int]bool)
				for _, step := range steps {
					x, y := i+step[0], j+step[1]
					if x >= 0 && y >= 0 && x < n && y < n && !visited[grid[x][y]] {
						area += cache[grid[x][y]]
						visited[grid[x][y]] = true
					}
				}
				size = max(size, area)
			}
		}
	}
	return size
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
