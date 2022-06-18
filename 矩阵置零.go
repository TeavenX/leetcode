package main

import "fmt"

func main() {
	//matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	r, c := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				r |= 1 << i
				c |= 1 << j
			}
		}
	}
	for row := 0; r > 0; row++ {
		if r&1 == 1 {
			for i := 0; i < m; i++ {
				matrix[row][i] = 0
			}
		}
		r >>= 1
	}
	for column := 0; c > 0; column++ {
		if c&1 == 1 {
			for i := 0; i < n; i++ {
				matrix[i][column] = 0
			}
		}
		c >>= 1
	}
}
