package main

func main() {

}

var steps = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func closedIsland(grid [][]int) int {
	ans := 0
	n, m := len(grid), len(grid[0])
	valid := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < m
	}
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if !valid(x, y) {
			return false
		}
		if grid[x][y] >= 1 {
			return true
		}
		grid[x][y] = 2
		for _, s := range steps {
			if !dfs(x+s[0], y+s[1]) {
				grid[x][y] = -1
				return false
			}
		}
		return true
	}
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] == 0 && dfs(i, j) {
				ans++
			}
		}
	}
	return ans
}
