package main

import "fmt"

func main() {
	matrix := [][]int{{1}, {3}}
	target := 1
	matrix = [][]int{{1}}
	target = 2
	matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target = 3
	fmt.Println(searchMatrixV3(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1
	for i < n {
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] == target {
			return true
		} else {
			break
		}
	}
	if i == n {
		return false
	}
	for j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		j--
	}
	return false
}

func searchMatrixV2(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1
	for i < n {
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] == target {
			return true
		} else {
			break
		}
	}
	if i == n {
		return false
	}
	left, right := 0, j
	for left <= right {
		mid := left + (right-left)/2
		if matrix[i][mid] == target {
			return true
		} else if matrix[i][mid] > target {
			right = mid - 1
		} else if matrix[i][mid] < target {
			left = mid + 1
		}
	}
	return false
}

func searchMatrixV3(matrix [][]int, target int) bool {
	n := len(matrix)
	i, j := 0, len(matrix[0])-1
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[mid][j] == target {
			return true
		} else if matrix[mid][j] > target {
			i = mid
			right = mid - 1
		} else if matrix[mid][j] < target {
			i = mid + 1
			left = mid + 1
		}
	}
	if i == n {
		return false
	}
	left, right = 0, j
	for left <= right {
		mid := left + (right-left)/2
		if matrix[i][mid] == target {
			return true
		} else if matrix[i][mid] > target {
			right = mid - 1
		} else if matrix[i][mid] < target {
			left = mid + 1
		}
	}
	return false
}
