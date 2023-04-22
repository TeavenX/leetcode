package main

func main() {

}

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	dp := make([][1002]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			d := nums[i] - nums[j]
			if d < 0 {
				d = 500 - d
			}
			dp[i][d] = max(dp[i][d], dp[j][d]+1)
			ans = max(ans, dp[i][d])
		}
	}
	return ans + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
