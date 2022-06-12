package main

import "sort"

func main() {

}

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)
	sort.Ints(nums)
	var traceback func()
	traceback = func() {
		if len(path) == n {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			num := nums[i]
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, num)
			used[i] = true
			traceback()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	traceback()
	return result
}
