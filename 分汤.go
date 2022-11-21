package main

import "fmt"

func main() {
	for i := 1; i < 1e9; i++ {
		r := soupServings(i)
		if r > 1-1e-5 {
			fmt.Println(i)
			// 4451
			break
		}
	}
}

func soupServings(n int) float64 {
	if n >= 4451 {
		return 1
	}
	n = (n + 24) / 25
	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
		dp[0][i] = 1
		if i == 0 {
			dp[0][i] = 0.5
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = 0.25 * (dp[mtz(i-4)][j] + dp[mtz(i-3)][mtz(j-1)] + dp[mtz(i-2)][mtz(j-2)] + dp[mtz(i-1)][mtz(j-3)])
		}
	}
	return dp[n][n]
}

func mtz(a int) int {
	if a < 0 {
		return 0
	}
	return a
}
