package main

import "fmt"

func main() {
	//costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	costs := [][]int{{3, 5, 3}, {6, 17, 6}, {7, 13, 18}, {9, 10, 18}}
	fmt.Println(minCostV4(costs))
}

func minCost(costs [][]int) int {
	dp := make([][]int, len(costs))
	for i, cost := range costs {
		dp[i] = make([]int, 3)
		for j := range cost {
			if i == 0 {
				dp[i][j] = cost[j]
			} else {
				dp[i][j] = min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3]) + cost[j]
			}
		}
	}
	return min(dp[len(costs)-1][0], min(dp[len(costs)-1][1], dp[len(costs)-1][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostV2(costs [][]int) int {
	dp := [2][3]int{}
	for j, cost := range costs[0] {
		dp[0][j] = cost
	}
	for _, cost := range costs[1:] {
		for j := range cost {
			dp[1][j] = min(dp[0][(j+1)%3], dp[0][(j+2)%3]) + cost[j]
		}
		dp[0] = dp[1]
	}
	return min(dp[0][0], min(dp[0][1], dp[0][2]))
}

func minCostV3(costs [][]int) int {
	preDP := make([]int, 3)
	curDP := make([]int, 3)
	for j, cost := range costs[0] {
		preDP[j] = cost
	}
	for _, cost := range costs[1:] {
		for j := range cost {
			curDP[j] = min(preDP[(j+1)%3], preDP[(j+2)%3]) + cost[j]
		}
		for j := range cost {
			preDP[j] = curDP[j]
		}
	}
	return min(preDP[0], min(preDP[1], preDP[2]))
}

func minCostV4(costs [][]int) int {
	n := len(costs)
	for i := 1; i < n; i++ {
		costs[i][0] = min(costs[i-1][1], costs[i-1][2]) + costs[i][0]
		costs[i][1] = min(costs[i-1][0], costs[i-1][2]) + costs[i][1]
		costs[i][2] = min(costs[i-1][1], costs[i-1][0]) + costs[i][2]
	}
	return min(costs[n-1][0], min(costs[n-1][1], costs[n-1][2]))
}
