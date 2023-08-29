package main

import "sort"

func main() {

}

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	has := make(map[int]bool)
	for _, x := range arr {
		has[x] = true
	}
	cache := make(map[int]int)
	var dfs func(val int) int
	dfs = func(val int) int {
		if v := cache[val]; v > 0 {
			return v
		}
		res := 1
		for _, x := range arr {
			if x >= val {
				break
			}
			if val%x == 0 && has[val/x] {
				res += dfs(x) * dfs(val/x)
			}
		}
		cache[val] = res
		return res
	}
	ans := 0
	for _, x := range arr {
		ans += dfs(x)
	}
	return ans % (1e9 + 7)
}
