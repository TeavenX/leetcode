package main

import "fmt"

func main() {
	grid := [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}
	result := shiftGrid(grid, 4)

	grid = [][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}
	result = shiftGrid(grid, 23)
	for i := 0; i < len(grid); i++ {
		fmt.Println(result[i])
	}
}

func shiftGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := (j + k) % m
			y := (i + (j+k)/m) % n
			result[y][x] = grid[i][j]
		}
	}
	return result
}
