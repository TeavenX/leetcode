package main

import "fmt"

func main() {
	//fmt.Println(minDistance("house", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	if word1 == "" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i-1][j]+1, dp[i][j-1]+1))
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
