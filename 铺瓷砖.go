package main

import "math"

func main() {

}

func tilingRectangle(n int, m int) int {
	cache := make([][]bool, n)
	for i := range cache {
		cache[i] = make([]bool, m)
	}

	valid := func(x, y, k int) bool {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if cache[x+i][y+j] {
					return false
				}
			}
		}
		return true
	}

	fill := func(x, y, k int, v bool) {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				cache[x+i][y+j] = v
			}
		}
	}

	ans := math.MaxInt

	var dfs func(x, y, count int)
	dfs = func(x, y, count int) {
		if count > ans {
			return
		}
		if x >= n {
			ans = count
			return
		}
		if y >= m {
			dfs(x+1, 0, count)
			return
		}
		if cache[x][y] {
			dfs(x, y+1, count)
			return
		}
		for k := min(n-x, m-y); k >= 1 && valid(x, y, k); k-- {
			fill(x, y, k, true)
			dfs(x, y+k, count+1)
			fill(x, y, k, false)
		}
	}
	dfs(0, 0, 0)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
