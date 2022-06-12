package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	nums = []int{4, 4, 4, 1, 4}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{{}}
	path := make([]int, 0)
	n := len(nums)
	var traceback func(int, int)
	sort.Ints(nums)
	traceback = func(startIdx, k int) {
		if k-1 == n {
			return
		}
		cache := make(map[int]bool)
		for i := startIdx; i < n; i++ {
			num := nums[i]
			if exists, _ := cache[num]; exists {
				continue
			}
			cache[num] = true
			path = append(path, num)
			if len(path) == k {
				temp := make([]int, k)
				copy(temp, path)
				result = append(result, temp)
				traceback(i+1, k+1)
				path = path[:len(path)-1]
			}
		}
	}
	traceback(0, 1)
	return result
}

func subsetsWithDupV2(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		ans = append(ans, append([]int{}, track...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return ans
}
