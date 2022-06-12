package main

import "fmt"

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	deep := len(triangle)
	if deep == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, deep)
	dp[0] = []int{triangle[0][0]}
	for i := 1; i < deep; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j <= i; j++ {
			if j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}
	result := dp[deep-1][0]
	for i := 1; i < len(dp[deep-1]); i++ {
		result = min(result, dp[deep-1][i])
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
