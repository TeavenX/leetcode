package main

import "math"

func main() {

}

func minSideJumps(obstacles []int) int {
	dp := [3]int{1, 0, 1}
	for _, ob := range obstacles[1:] {
		minCnt := math.MaxInt >> 1
		for i := 0; i < 3; i++ {
			if ob == i+1 {
				dp[i] = math.MaxInt >> 1
			} else {
				minCnt = min(minCnt, dp[i])
			}
		}
		for i := 0; i < 3; i++ {
			if ob != i+1 {
				dp[i] = min(dp[i], minCnt+1)
			}
		}
	}
	return min(min(dp[0], dp[1]), dp[2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
