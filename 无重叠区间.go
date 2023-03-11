package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	// intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {-100, -2}, {5, 7}}
	fmt.Println(eraseOverlapIntervals(intervals))
}

type pair struct {
	i, end int
}

func eraseOverlapIntervalsV1TLE(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1] {
			return true
		}
		return false
	})
	n := len(intervals)
	cache := make(map[pair]int)
	var dfs func(i int, end int) int
	dfs = func(i int, end int) int {
		if i == n {
			return 0
		}
		p := pair{i, end}
		if c, exist := cache[p]; exist {
			return c
		}
		t := dfs(i+1, end)
		if end <= intervals[i][0] {
			t = max(t, dfs(i+1, intervals[i][1])+1)
		}
		cache[p] = t
		return t
	}
	ans := dfs(0, math.MinInt)
	return n - ans
}

func eraseOverlapIntervalsV2TLE(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1] {
			return true
		}
		return false
	})
	n := len(intervals)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	return n - ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	ans := 1
	for i := 1; i < n; i++ {
		if intervals[i][0] >= intervals[i-1][1] {
			ans++
		} else {
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
		}
	}
	return n - ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
