package main

func main() {

}

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	dp[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	result := n - dp[n][m]
	result += m - dp[n][m]
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
