package main

import "sort"

func main() {

}

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	nums = append(nums, 0)
	for i := range nums[:len(nums)-1] {
		i++
		nums[i] += nums[i-1]
	}
	ans := make([]int, len(queries))
	for i, s := range queries {
		for j, num := range nums[:len(nums)-1] {
			if num <= s {
				ans[i] = j + 1
			}
		}
	}
	return ans
}

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	n := len(nums)
	arr := make([]int, n)
	arr[0] = nums[0]
	for i := 1; i < n; i++ {
		arr[i] = arr[i-1] + nums[i]
	}
	ans := make([]int, len(queries))
	for i, s := range queries {
		ans[i] = sort.SearchInts(arr, s+1)
	}
	return ans
}
