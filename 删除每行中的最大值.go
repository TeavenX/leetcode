package main

import "sort"

func main() {

}

func deleteGreatestValue(grid [][]int) int {
	ans := 0
	for i := range grid {
		sort.Ints(grid[i])
	}
	for j := 0; j < len(grid[0]); j++ {
		res := grid[0][j]
		for i := 1; i < len(grid); i++ {
			res = max(res, grid[i][j])
		}
		ans += res
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
