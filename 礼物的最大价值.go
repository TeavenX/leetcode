package main

func main() {

}

func maxValue(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	var traceback func(i, j int) int
	traceback = func(i, j int) int {
		if i >= n || j >= m {
			return 0
		}
		if i == n-1 && j == m-1 {
			return grid[i][j]
		}
		if grid[i][j] < 0 {
			return -grid[i][j]
		}
		t := grid[i][j] + max(traceback(i+1, j), traceback(i, j+1))
		grid[i][j] = -t
		return t
	}
	ans := traceback(0, 0)
	return ans
}

func maxValueV2(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	for j := 1; j < m; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < n; i++ {
		grid[i][0] += grid[i-1][0]
		for j := 1; j < m; j++ {
			grid[i][j] += max(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[n-1][m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
