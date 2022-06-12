package main

import "sort"

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	result := make([][]int, 0)
	sort.Ints(candidates)
	if candidates[0] > target {
		return result
	}
	path := make([]int, 0)
	sum := 0
	var traceback func()
	traceback = func() {
		if sum == target {
			result = append(result, append([]int{}, path...))
			return
		} else if sum > target {
			return
		}
		for i := 0; i < n; i++ {
			num := candidates[i]
			if len(path) > 0 && num < path[len(path)-1] {
				continue
			}
			path = append(path, num)
			sum += num
			traceback()
			sum -= num
			path = path[:len(path)-1]
		}
	}
	traceback()
	return result
}
