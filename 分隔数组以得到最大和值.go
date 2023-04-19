package main

func main() {

}

func maxSumAfterPartitioning(arr []int, k int) int {
	cache := make([]int, len(arr))
	for i := range cache {
		cache[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if cache[i] > 0 {
			return cache[i]
		}
		mx, res := 0, 0
		for j := i; j >= 0 && j > i-k; j-- {
			mx = max(mx, arr[j])
			res = max(res, (i-j+1)*mx+dfs(j-1))
		}
		cache[i] = res
		return res
	}
	return dfs(len(arr) - 1)
}

func maxSumAfterPartitioningV2(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		mx := 0
		for j := i; j >= 0 && j > i-k; j-- {
			mx = max(mx, arr[j])
			dp[i+1] = max(dp[i+1], dp[j]+(i-j+1)*mx)
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
