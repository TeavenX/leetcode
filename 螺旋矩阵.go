package main

import (
	"fmt"
)

func main() {
	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrderError(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	var bfs func(i, j, nt, nb, ml, mr int)
	result := make([]int, 0)
	bfs = func(i, j, nt, nb, ml, mr int) {
		if i > nb-1 || i < nt || j > mr-1 || j < ml {
			return
		}
		for j < mr {
			result = append(result, matrix[i][j])
			j++
		}
		for i < nb {
			result = append(result, matrix[i][j])
			i++
		}
		for j > ml {
			result = append(result, matrix[i][j])
			j--
		}
		for i > nt {
			result = append(result, matrix[i][j])
			i--
		}
		bfs(i+1, j+1, nt+1, nb-1, ml+1, mr-1)
	}
	bfs(0, 0, 0, n-1, 0, m-1)
	return result
}

func spiralOrder(matrix [][]int) []int {
	rows, columns := len(matrix), len(matrix[0])
	idx, left, top, bottom, right := 0, 0, 0, rows-1, columns-1
	result := make([]int, rows*columns)
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			result[idx] = matrix[top][column]
			idx++
		}
		for row := top + 1; row <= bottom; row++ {
			result[idx] = matrix[row][right]
			idx++
		}
		if left < right && top < bottom {
			for column := right - 1; column >= left; column-- {
				result[idx] = matrix[bottom][column]
				idx++
			}
			for row := bottom - 1; row > top; row-- {
				result[idx] = matrix[row][left]
				idx++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
