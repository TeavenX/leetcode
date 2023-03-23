package main

import (
	"fmt"
	"sort"
)

func main() {
	scores := []int{1, 2, 3, 5}
	ages := []int{8, 9, 10, 1}
	fmt.Println(bestTeamScore(scores, ages))
}

func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)
	index := make([]int, n)
	for i := range index {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		i, j = index[i], index[j]
		return ages[i] < ages[j] || ages[i] == ages[j] && scores[i] < scores[j]
	})
	cache := make([]int, n)
	ans := 0
	for i, idx := range index {
		fmt.Println(ages[idx])
		res := 0
		for j := 0; j < i; j++ {
			jdx := index[j]
			if scores[jdx] <= scores[idx] {
				res = max(res, cache[j])
			}
		}
		cache[i] = res + scores[idx]
		ans = max(ans, cache[i])
	}
	fmt.Println(cache)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
