package main

func main() {
	champagneTower(25, 6, 1)
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
	if poured == 0 {
		return 0
	}
	if poured >= 100 {
		return 1
	}
	dp := make([][]float64, query_row+1)
	dp[0] = []float64{float64(poured)}
	for i := 1; i <= query_row; i++ {
		dp[i] = make([]float64, i+1)
		dp[i][0] = moreThanZero(dp[i-1][0]-1) / 2
		for j := 1; j < i; j++ {
			dp[i][j] = moreThanZero(dp[i-1][j]-1) / 2
			dp[i][j] += moreThanZero(dp[i-1][j-1]-1) / 2
		}
		dp[i][i] = moreThanZero(dp[i-1][i-1]-1) / 2
	}
	if dp[query_row][query_glass] > 1 {
		dp[query_row][query_glass] = 1
	}
	return dp[query_row][query_glass]
}

func moreThanZero(a float64) float64 {
	if a < 0 {
		return 0
	}
	return a
}
