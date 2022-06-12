package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2V2(candidates, target))
}

func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	result := make([][]int, 0)
	sort.Ints(candidates)
	path := make([]int, 0)
	sum := 0
	used := make(map[int]bool)
	var traceback func(startIdx int)
	traceback = func(startIdx int) {
		if sum == target {
			result = append(result, append([]int{}, path...))
			return
		} else if sum >= target {
			return
		}
		for i := startIdx; i < n; i++ {
			num := candidates[i]
			if i > 0 && candidates[i-1] == num && !used[i-1] {
				continue
			}
			path = append(path, num)
			sum += num
			used[i] = true
			traceback(i + 1)
			path = path[:len(path)-1]
			sum -= num
			used[i] = false
		}
	}
	traceback(0)
	return result
}

func combinationSum2V2(candidates []int, target int) [][]int {
	n := len(candidates)
	result := make([][]int, 0)
	sort.Ints(candidates)
	path := make([]int, 0)
	sum := 0
	var traceback func(startIdx int)
	traceback = func(startIdx int) {
		//fmt.Println("第", startIdx+1, "层回溯--/")
		if sum == target {
			result = append(result, append([]int{}, path...))
			return
		} else if sum >= target {
			return
		}
		for i := startIdx; i < n; i++ {
			num := candidates[i]
			//fmt.Println("i:[", i, "], num:[", num, "]")
			if i > startIdx && candidates[i-1] == num {
				continue
			}
			path = append(path, num)
			sum += num
			traceback(i + 1)
			path = path[:len(path)-1]
			sum -= num
		}
	}
	traceback(0)
	return result
}
