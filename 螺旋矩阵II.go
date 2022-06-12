package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	for num := 1; num <= n*n; {
		for j := left; j <= right; j++ {
			result[top][j] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			result[i][right] = num
			num++
		}
		right--
		for j := right; j >= left; j-- {
			result[bottom][j] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			result[i][left] = num
			num++
		}
		left++
	}
	return result
}
