package main

func main() {

}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	mouseMoveCount := 0
	var mouseX, mouseY, catX, catY, foodX, foodY int
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'C' {
				catX, catY = i, j
			}
			if grid[i][j] == 'M' {
				mouseX, mouseY = i, j
			}
			if grid[i][j] == 'F' {
				foodX, foodY = i, j
			}
		}
	}
	for mouseMoveCount < 1000 {

	}
}
