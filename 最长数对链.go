package main

import (
	"fmt"
	"sort"
)

func main() {
	pairs := [][]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println(findLongestChain(pairs))
}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	count := 1
	cur := pairs[0][1]
	for _, pair := range pairs {
		if cur < pair[0] {
			cur = pair[1]
			count++
		}
	}
	return count
}
