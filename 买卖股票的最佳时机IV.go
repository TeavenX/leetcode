package main

import "math"

func main() {

}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	// 因为交易日一共只有len(prices)天，买卖股票不会在同一天，因此如果k >= len(prices)/2，可以认为k是无穷大
	// 此时可以简化下面的遍历流程，不需要考虑k
	dp := make([][][2]int, n+1)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}
	for i := 1; i <= k; i++ {
		dp[0][i][1] = math.MinInt // 买入之后收益可能会是负数，这里提前初始化为最小值，避免后面max不生效
	}
	for i, price := range prices {
		for j := k; j > 0; j-- {
			dp[i+1][j][0] = max(dp[i][j][0], dp[i][j][1]+price)   // 卖出
			dp[i+1][j][1] = max(dp[i][j][1], dp[i][j-1][0]-price) // 买入
		}
	}
	return dp[n][k][0] // 买入过k次，全部卖出后的收益
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
