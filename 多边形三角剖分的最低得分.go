package main

import (
	"fmt"
	"math"
)

func main() {
	values := []int{1, 3, 1, 4, 1, 5}
	fmt.Println(minScoreTriangulation(values))
}

func minScoreTriangulationTLE(values []int) int {
	n := len(values)
	if n < 3 {
		return math.MaxInt
	} else if n == 3 {
		return values[0] * values[1] * values[2]
	}
	ans := math.MaxInt
	for i := 0; i < n-2; i++ {
		for j := i + 2; (i > 0 && j < n) || j < n-1; j++ {
			ans = min(ans, minScoreTriangulationTLE(values[i:j+1])+minScoreTriangulationTLE(append(append([]int{}, values[j:n]...), values[0:i+1]...)))
		}
	}
	return ans
}

func minScoreTriangulation(values []int) int {
	n := len(values)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (ans int) {
		if i+1 == j {
			return 0
		}
		if cache[i][j] > 0 {
			return cache[i][j]
		}
		defer func() {
			cache[i][j] = ans
		}()
		ans = 5e10
		for k := i + 1; k < j; k++ {
			res := dfs(i, k) + dfs(k, j) + values[i]*values[j]*values[k]
			ans = min(ans, res)
		}
		return ans
	}
	return dfs(0, n-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
