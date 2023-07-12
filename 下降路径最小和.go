package main

import "math"

func main() {

}

var steps = [][]int{{1, -1}, {1, 0}, {1, 1}}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1e5
		}
	}
	valid := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < n
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i == n {
			return 0
		}
		if !valid(i, j) {
			return 1e6
		}
		if v := cache[i][j]; v > -1e5 {
			return v
		}
		defer func() {
			cache[i][j] = res
		}()
		res = 1e5
		for _, s := range steps {
			res = min(res, dfs(i+s[0], j+s[1]))
		}
		return res + matrix[i][j]
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, dfs(0, i))
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var steps = [][]int{{1, -1}, {1, 0}, {1, 1}}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	valid := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < n
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i == n {
			return 0
		}
		if !valid(i, j) {
			return 1e6
		}
		if v := matrix[i][j]; v > 100 || v < -100 {
			t := abs(v) - 101
			if v < 0 {
				return -t
			}
			return t
		}
		defer func() {
			t := abs(res) + 101
			matrix[i][j] = t
			if res < 0 {
				matrix[i][j] = -t
			}
		}()
		res = 1e5
		for _, s := range steps {
			res = min(res, dfs(i+s[0], j+s[1]))
		}
		return res + matrix[i][j]
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, dfs(0, i))
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
