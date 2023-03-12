package main

func main() {

}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	// dp[i][k][0], dp[i][k][1]
	// i表示第i天，k表示交易上限
	// 0 没有买，1买了
	// dp[i][0] = max(dp[i][0], dp[i][1] + prices[i]) // 卖股票卖了prices[i]
	// dp[i][1] = max(dp[i][1], dp[i][0] - prices[i]) // 买股票花了prices[i]
	dp[0][1] = -prices[0]
	for i, price := range prices {
		if i > 0 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price)
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-price)
		}
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
