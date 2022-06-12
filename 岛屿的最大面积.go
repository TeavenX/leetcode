package main

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			newArea := getArea(i, j, &grid)
			if newArea > maxArea {
				maxArea = newArea
			}
		}
	}
	return maxArea
}

func getArea(x, y int, grid *[][]int) int {
	area := 0
	if len(*grid) > x && x >= 0 && len((*grid)[0]) > y && y >= 0 {
		if (*grid)[x][y] == 1 {
			(*grid)[x][y] = 0
			area = 1
			area += getArea(x-1, y, grid)
			area += getArea(x+1, y, grid)
			area += getArea(x, y-1, grid)
			area += getArea(x, y+1, grid)
		}
	}
	return area
}
