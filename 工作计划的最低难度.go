package main

import "math"

func main() {

}

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}
	cache := make([][]int, d)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if v := cache[i][j]; v >= 0 {
			return v
		}
		defer func() {
			cache[i][j] = res
		}()
		if i == 0 {
			for _, v := range jobDifficulty[:j+1] {
				res = max(res, v)
			}
			return res
		}
		res = math.MaxInt
		mx := 0
		for k := j; k >= i; k-- {
			mx = max(mx, jobDifficulty[k])
			res = min(res, dfs(i-1, k-1)+mx)
		}
		return res
	}
	return dfs(d-1, n-1)
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
