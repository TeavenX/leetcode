package main

func main() {

}

var steps = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func uniquePathsIII(grid [][]int) int {
	sx, sy := 0, 0
	total := 0
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				sx, sy = i, j
			} else if v == 0 {
				total++
			}
		}
	}
	n, m := len(grid), len(grid[0])
	ans := 0
	var dfs func(i, j, left int)
	dfs = func(i, j, left int) {
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		v := grid[i][j]
		if v == -1 {
			return
		} else if v == 2 {
			if left == 0 {
				ans++
			}
			return
		}
		grid[i][j] = -1
		for _, step := range steps {
			dfs(i+step[0], j+step[1], left-1)
		}
		grid[i][j] = v
	}
	dfs(sx, sy, total+1)
	return ans
}
