package main

func main() {

}

var steps = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func numIslands(grid [][]byte) int {
	result := 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				result++
				walk(&grid, i, j, n, m)
			}
		}
	}
	return result
}

func walk(grid *[][]byte, i, j, n, m int) {
	(*grid)[i][j] = '0'
	for _, step := range steps {
		newi, newj := i+step[0], j+step[1]
		if newi < n && newi >= 0 && newj < m && newj >= 0 && (*grid)[newi][newj] == '1' {
			walk(grid, newi, newj, n, m)
		}
	}
}

func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	count := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= n || j >= m {
			return
		}
		if grid[i][j] == '1' {
			grid[i][j] = '0'
			for _, step := range steps {
				dfs(i+step[0], j+step[1])
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}
