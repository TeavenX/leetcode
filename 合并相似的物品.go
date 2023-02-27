package main

import "sort"

func main() {

}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	cache := make(map[int]int)
	for _, item := range items1 {
		cache[item[0]] += item[1]
	}
	for _, item := range items2 {
		cache[item[0]] += item[1]
	}
	ans := make([][]int, 0, len(cache))
	for k, v := range cache {
		ans = append(ans, []int{k, v})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}
