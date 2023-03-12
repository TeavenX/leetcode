package main

import "math"

func main() {

}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][3][2]int, n+1)
	// dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + price) // 第i天，已经买入 2 次之后，本次 卖出 所能获得的最大利润
	// dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - price) // 第i天，已经买入 1 次之后，本次 买入 所能获得的最大利润
	// dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + price) // 第i天，已经买入 1 次之后，本次 卖出 所能获得的最大利润
	// dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - price) // 第i天，已经买入 0 次之后，本次 买入 所能获得的最大利润

	// dp[0][1][1] = math.MinInt
	// dp[0][2][1] = math.MinInt
	// for i, price := range prices {
	//     dp[i+1][2][0] = max(dp[i][2][0], dp[i][2][1] + price)
	//     dp[i+1][2][1] = max(dp[i][2][1], dp[i][1][0] - price)
	//     dp[i+1][1][0] = max(dp[i][1][0], dp[i][1][1] + price)
	//     dp[i+1][1][1] = max(dp[i][1][1], dp[i][0][0] - price)
	// }
	for i := 1; i < 3; i++ {
		dp[0][i][1] = math.MinInt
	}
	for i, price := range prices {
		for j := 2; j > 0; j-- {
			dp[i+1][j][0] = max(dp[i][j][0], dp[i][j][1]+price)
			dp[i+1][j][1] = max(dp[i][j][1], dp[i][j-1][0]-price)
		}
	}
	return dp[n][2][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
