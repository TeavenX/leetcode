package main

import (
	"fmt"
	"sort"
)

func main() {
	matchsticks := []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}
	fmt.Println(makesquare(matchsticks))
}

func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	if n < 4 {
		return false
	}
	totalLen := 0
	for _, stickLen := range matchsticks {
		totalLen += stickLen
	}
	if totalLen%4 != 0 {
		return false
	}
	perLen := totalLen >> 2
	sides := make([]int, 4)
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == n {
			return true
		}
		for i := 0; i < 4; i++ {
			sides[i] += matchsticks[idx]
			if sides[i] <= perLen && dfs(idx+1) {
				return true
			}
			sides[i] -= matchsticks[idx]
		}
		return false
	}
	return dfs(0)
}
