package main

func main() {

}

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	cache := make(map[int]int)
	dp := make([][]int, n)
	for idx, val := range arr {
		dp[idx] = make([]int, n)
		cache[val] = idx
	}
	maxN := 0
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			diff := arr[j] - arr[i]
			if idx, ok := cache[diff]; ok {
				if idx < i {
					dp[i][j] = max(3, dp[idx][i]+1)
					maxN = max(maxN, dp[i][j])
				}
			}
		}
	}
	return maxN
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
