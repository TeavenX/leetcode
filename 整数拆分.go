package main

func main() {

}

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreakV2(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	n3 := n / 3
	r3 := n - n3*3
	if r3 == 1 {
		n3--
		r3 += 3
	}
	result := 1
	for i := 0; i < n3; i++ {
		result *= 3
	}
	if r3 != 0 {
		result *= r3
	}
	return result
}
