package main

import "math"

func main() {

}

func maximumSum(arr []int) int {
	// dfs(i, j) 其中i表示第i个数，j表示删除（1）和不删除（0）
	// dfs(i, 0) = dfs(i-1, 0) + arr[i]
	// dfs(i, 1) = max(dfs(i-1, 1) + arr[i], dfs(i-1, 0))
	n := len(arr)
	cache := make([][2]int, n)
	for i := range cache {
		cache[i][0] = math.MinInt
		cache[i][1] = math.MinInt
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i == -1 {
			return math.MinInt / 2
		}
		if v := cache[i][j]; v > math.MinInt {
			return v
		}
		defer func() {
			cache[i][j] = res
		}()
		if j == 0 {
			return max(dfs(i-1, 0), 0) + arr[i]
		}
		return max(dfs(i-1, 1)+arr[i], dfs(i-1, 0))
	}
	ans := math.MinInt
	for i := n - 1; i >= 0; i-- {
		ans = max(ans, max(dfs(i, 1), dfs(i, 0)))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumSum(arr []int) int {
	n := len(arr)
	f := make([][2]int, n+1)
	for i := range f {
		f[i] = [2]int{math.MinInt / 2, math.MinInt / 2}
	}
	ans := math.MinInt
	for i := 0; i < n; i++ {
		f[i+1][0] = max(f[i][0], 0) + arr[i]
		f[i+1][1] = max(f[i][1]+arr[i], f[i][0])
		ans = max(ans, max(f[i+1][0], f[i+1][1]))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
