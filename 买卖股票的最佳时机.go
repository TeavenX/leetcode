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

func maxProfit20220702(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}
	dp := make([]int, n)
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], prices[i]-minPrice)
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return dp[n-1]
}

type singleStack struct {
	stack []int
	min   int
}

func (s *singleStack) push(val int) int {
	n := len(s.stack)
	maxDist := 0
	if n > 0 {
		maxDist = s.stack[n-1] - s.min
	}
	if val < s.min {
		s.stack = []int{val}
		s.min = val
	} else if n == 0 || s.stack[n-1] < val {
		if n == 0 {
			s.min = val
		}
		s.stack = append(s.stack, val)
	}
	return maxDist
}

func maxProfitSingleStack(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}
	stack := singleStack{}
	maxVal := 0
	for _, price := range prices {
		maxVal = max(maxVal, stack.push(price))
	}
	return max(maxVal, stack.stack[len(stack.stack)-1]-stack.min)
}
