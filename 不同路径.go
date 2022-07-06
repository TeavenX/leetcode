package main

import "math/big"

func main() {

}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 0
			} else if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func uniquePaths(m int, n int) int {
	// 因为m*n的方格内，机器人一定会走m+n-2步（m+n是边长，减去起始位置和mn重合的位置）
	// 其中 m-1 步向下移动， n-1 步向右移动
	// 所以这里只要求 m+n-2 和 n-1 的排列组合
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
