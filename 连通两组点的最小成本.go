package main

import "math"

func main() {

}

func connectTwoGroups(cost [][]int) int {
	n, m := len(cost), len(cost[0])
	minCost := make([]int, m)
	for i := 0; i < m; i++ {
		minCost[i] = math.MaxInt
		for _, c := range cost {
			minCost[i] = min(minCost[i], c[i])
		}
	}
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, 1<<m)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			for i := 0; i < m; i++ {
				if j>>i&1 == 1 {
					res += minCost[i]
				}
			}
			return res
		}
		v := &cache[i][j]
		if *v >= 0 {
			return *v
		}
		res = math.MaxInt
		for k := 0; k < m; k++ {
			res = min(res, dfs(i-1, j&^(1<<k))+cost[i][k])
		}
		*v = res
		return res
	}
	return dfs(n-1, 1<<m-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
