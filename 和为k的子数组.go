package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}

func subarraySum(nums []int, k int) int {
	cache := map[int]int{0: 1}
	result := 0
	preSum := 0
	for _, num := range nums {
		preSum += num
		cnt := cache[preSum-k]
		result += cnt
		cache[preSum]++
	}
	return result
}
