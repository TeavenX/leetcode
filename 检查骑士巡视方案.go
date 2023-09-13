package main

func main() {

}

var steps = [][]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	n := len(grid)
	end := n*n - 1
	var dfs func(x, y, idx int) bool
	dfs = func(x, y, idx int) bool {
		if idx == end {
			return true
		}
		for _, step := range steps {
			i, j := x+step[0], y+step[1]
			if i < 0 || j < 0 || i >= n || j >= n || grid[i][j] != idx+1 {
				continue
			}
			if grid[i][j] == idx+1 {
				return dfs(i, j, idx+1)
			}
		}
		return false
	}
	return dfs(0, 0, 0)
}
