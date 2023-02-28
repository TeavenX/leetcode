package main

func main() {

}

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			for k := 0; k < 3; k++ {
				for q := 0; q < 3; q++ {
					grid[i][j] = max(grid[i][j], grid[i+k][j+q])
				}
			}
		}
		grid[i] = grid[i][:n-2]
	}
	return grid[:n-2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
