package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"IEO", "Sgizfdfrims", "QTASHKQ", "Vk", "RPJOFYZUBFSIYp", "EPCFFt", "VOYGWWNCf", "WSpmqvb"}
	heights := []int{17233, 32521, 14087, 42738, 46669, 65662, 43204, 8224}
	fmt.Println(sortPeople(names, heights))
}

func sortPeople(names []string, heights []int) []string {
	n := len(names)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	ans := make([]string, n)
	sort.Slice(idx, func(i, j int) bool {
		return heights[idx[i]] > heights[idx[j]]
	})
	for i := range idx {
		ans[i] = names[idx[i]]
	}
	return ans
}
