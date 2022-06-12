package main

func main() {

}

func projectionArea(grid [][]int) int {
	n := len(grid[0])
	sum := 0
	for i := 0; i < n; i++ {
		horizonMax := grid[i][0]
		verticalMax := grid[0][i]
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				sum += 1
			}
			if grid[i][j] > horizonMax {
				horizonMax = grid[i][j]
			}
			if grid[j][i] > verticalMax {
				verticalMax = grid[j][i]
			}
		}
		sum += horizonMax
		sum += verticalMax
	}
	return sum
}
