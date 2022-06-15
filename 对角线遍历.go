package main

import "fmt"

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(findDiagonalOrder(mat))
}

func findDiagonalOrder(mat [][]int) []int {
	n := len(mat)
	if n == 1 {
		return mat[0]
	}
	m := len(mat[0])
	result := make([]int, m*n)
	if m == 1 {
		for i, matrix := range mat {
			result[i] = matrix[0]
		}
	}
	idx := 0
	flag := true
	for i := 0; i < n+m; i++ {
		curN, curM := n, m
		if !flag {
			curN, curM = m, n
		}
		x := min(i, curN-1)
		y := i - x
		for x >= 0 && y < curM {
			if !flag {
				result[idx] = mat[y][x]
			} else {
				result[idx] = mat[x][y]
			}
			x--
			y++
			idx++
		}
		flag = !flag
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
