package main

import (
	"fmt"
	"math"
)

func main() {
	// [[0,1,-1],[1,0,-1],[1,1,1]]
	// [[1,1,-1],[1,-1,1],[-1,1,1]]
	grid := make([][]int, 3)
	grid[0] = []int{0, 1, -1}
	grid[1] = []int{1, 0, -1}
	grid[2] = []int{1, 1, 1}
	cherryPickup(grid)
}

func cherryPickupError(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		if grid[i][0] == -1 || dp[i-1][0] == -1 {
			dp[i][0] = -1
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
			grid[i][0] = 0
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if i == 0 {
				if grid[0][j] == -1 || dp[0][j-1] == -1 {
					dp[0][j] = -1
				} else {
					dp[0][j] = dp[0][j-1] + grid[0][j]
					grid[0][j] = 0
				}
			} else {
				if grid[i][j] != -1 && (dp[i-1][j] != -1 || dp[i][j-1] != -1) {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
					grid[i][j] = 0
				} else {
					dp[i][j] = -1
				}
			}
		}
	}
	//for i := n - 2; i >= 0; i-- {
	//    if grid[i][n-1] == -1 || dp[i+1][n-1] == -1 {
	//        dp[i][n-1] = -1
	//    } else {
	//        dp[i][n-1] = dp[i+1][n-1] + grid[i][n-1]
	//        grid[i][n-1] = 0
	//    }
	//}
	//for i := n - 1; i >= 0; i-- {
	//    for j := n - 2; j >= 0; j-- {
	//        if i == n-1 {
	//            if grid[i][j] == -1 || dp[i][j+1] == -1 {
	//                dp[i][j] = -1
	//            } else {
	//                dp[i][j] = dp[i][j+1] + grid[i][j]
	//                grid[i][j] = 0
	//            }
	//        } else {
	//            if grid[i][j] != -1 && (dp[i+1][j] != -1 || dp[i][j+1] != -1) {
	//                dp[i][j] = max(dp[i+1][j], dp[i][j+1]) + grid[i][j]
	//                grid[i][j] = 0
	//            } else {
	//                dp[i][j] = -1
	//            }
	//        }
	//    }
	//}
	for _, v := range dp {
		fmt.Println(v)
	}
	for _, v := range grid {
		fmt.Println(v)
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MinInt
		}
	}
	dp[0][0] = grid[0][0]
	for k := 1; k < n*2-1; k++ {
		for x1 := min(k, n-1); x1 >= max(k-n+1, 0); x1-- {
			for x2 := min(k, n-1); x2 >= x1; x2-- {
				y1, y2 := k-x1, k-x2
				if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
					dp[x1][x2] = math.MinInt
					continue
				}
				res := dp[x1][x2]
				if x1 > 0 {
					res = max(res, dp[x1-1][x2])
				}
				if x2 > 0 {
					res = max(res, dp[x1][x2-1])
				}
				if x1 > 0 && x2 > 0 {
					res = max(res, dp[x1-1][x2-1])
				}
				res += grid[x1][y1]
				if x2 != x1 {
					res += grid[x2][y2]
				}
				dp[x1][x2] = res
			}
		}
	}
	return max(dp[n-1][n-1], 0)
}
