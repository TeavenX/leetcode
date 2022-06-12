package main

import "fmt"

func main() {
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	grid = [][]int{{0, 1, 1, 0, 0, 0}, {0, 1, 0, 1, 1, 0}, {0, 1, 1, 0, 1, 0}, {0, 0, 0, 1, 1, 0}, {1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 0}}
	fmt.Println(shortestPathBinaryMatrix(grid))
}

var steps = [][]int{{1, 1}, {1, 0}, {0, 1}, {1, -1}, {-1, 1}, {-1, -1}, {-1, 0}, {0, -1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}
	queue := [][]int{{0, 0}}
	grid[0][0] = 1
	for len(queue) > 0 {
		level := len(queue)
		for node := 0; node < level; node++ {
			i, j := queue[node][0], queue[node][1]
			for _, step := range steps {
				newi, newj := i+step[0], j+step[1]
				if 0 <= newi && newi < n && 0 <= newj && newj < n {
					if grid[newi][newj] == 0 {
						queue = append(queue, []int{newi, newj})
						grid[newi][newj] = grid[i][j] + 1
					}
					if newi == n-1 && newj == n-1 {
						return grid[newi][newj]
					}
				}
			}
		}
		queue = queue[level:]
	}
	return -1
}
