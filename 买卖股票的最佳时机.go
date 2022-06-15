package main

func main() {

}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}
	dp := make([]int, n)
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		price := prices[i]
		dp[i] = max(dp[i-1], price-minPrice)
		if price < minPrice {
			minPrice = price
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
