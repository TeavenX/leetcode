package main

import "fmt"

func main() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	cache := make([]int, n)
	cache[1] = min(cost[0], cost[1])
	for i := 2; i < n; i++ {
		cache[i] = min(cache[i-1]+cost[i], cache[i-2]+cache[i-1])
	}
	fmt.Println(cache)
	return cache[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
