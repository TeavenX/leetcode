package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxStrength([]int{8, 6, 0, 5, -4, -8, -4, 9, -1, 6, -4, 8, -5}))
}

func maxStrength(nums []int) int64 {
	n := len(nums)
	ans := math.MinInt
	cache := make(map[int]bool)
	var dfs func(sts int, state int)
	dfs = func(sts int, state int) {
		ans = max(ans, sts)
		if cache[state] {
			return
		}
		cache[state] = true
		for i := 0; i < n; i++ {
			if state>>i&1 == 0 {
				if sts == math.MinInt {
					sts = 1
				}
				dfs(sts*nums[i], state|1<<i)
			}
		}
		return
	}
	dfs(ans, 0)
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
