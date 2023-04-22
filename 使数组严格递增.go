package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr1 := []int{1, 6, 7, 6, 7}
	// arr2 := []int{2, 3, 7, 1, 5}
	arr1 := []int{5, 16, 19, 2, 1, 12, 7, 14, 5, 16}
	arr2 := []int{6, 17, 4, 3, 6, 13, 4, 3, 18, 17, 16, 7, 14, 1, 16}
	fmt.Println(makeArrayIncreasing(arr1, arr2))
}

func makeArrayIncreasingTLE(arr1 []int, arr2 []int) int {
	n := len(arr1)
	if n == 1 {
		return 0
	}
	sort.Ints(arr2)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n+1 {
			return 0
		}
		// 改成二分查找
		var res int = 2e4
		pre := arr1[i-1]
		cur := arr1[i]
		for _, num := range arr2 {
			if num > pre {
				arr1[i] = num
				res = min(res, 1+dfs(i+1))
				arr1[i] = cur
			}
		}
		if arr1[i] > arr1[i-1] {
			res = min(res, dfs(i+1))
		}
		return res
	}
	arr1 = append([]int{-1}, arr1...)
	ans := dfs(1)
	if ans == 2e4 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	n := len(arr1)
	if n == 1 {
		return 0
	}
	sort.Ints(arr2)
	cache := make(map[int]int)
	var dfs func(i, pre int) int
	dfs = func(i, pre int) int {
		if i == n {
			return 0
		}
		key := pre*1e4 + i
		if v, ok := cache[key]; ok {
			return v
		}
		var res int = 2e4
		if arr1[i] > pre {
			res = min(res, dfs(i+1, arr1[i]))
		}
		idx := sort.SearchInts(arr2, pre+1)
		if idx < len(arr2) {
			res = min(res, 1+dfs(i+1, arr2[idx]))
		}
		cache[key] = res
		return res
	}
	ans := dfs(0, -1)
	if ans == 2e4 {
		return -1
	}
	return ans
}
