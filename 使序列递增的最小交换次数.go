package main

func main() {

}

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] && nums2[i] > nums1[i-1] && nums1[i] > nums2[i-1] {
			dp[i][0] = min(dp[i-1][0], dp[i-1][1])
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		} else {
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
