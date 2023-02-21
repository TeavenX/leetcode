package main

import "math"

func main() {

}

func stoneGameII(piles []int) int {
	n := len(piles)
	for i := n - 2; i >= 0; i-- {
		piles[i] += piles[i+1]
	}
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, (n+1)>>2+1)
	}
	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i+m*2 >= n {
			return piles[i]
		}
		if cache[i][m] > 0 {
			return cache[i][m]
		}
		t := math.MaxInt
		for j := 1; j <= m*2; j++ {
			t = min(t, dfs(i+j, max(m, j)))
		}
		ans := piles[i] - t
		cache[i][m] = ans
		return ans
	}
	return dfs(0, 1)
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
