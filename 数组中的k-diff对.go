package main

import "fmt"

func main() {
	nums := []int{1, 3, 1, 5, 4}
	fmt.Println(findPairs(nums, 0))
}

func findPairs(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	result := 0
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}
	for num, v := range cache {
		if k == 0 {
			if v > 1 {
				result++
			}
		} else {
			if _, exist := cache[num+k]; exist {
				result++
			}
		}
	}
	return result
}
