package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 1, 5, 4}
	fmt.Println(findPairsDoublePtr(nums, 0))
}

func findPairs(nums []int, k int) int {
	if len(nums) == 1 {
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

func findPairsDoublePtr(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	sort.Ints(nums)
	result := 0
	j := 0
	for i, num := range nums {
		if i == 0 || num != nums[i-1] {
			for j < n && (nums[j] < num+k || j <= i) {
				j++
			}
			if j < n && nums[j] == num+k {
				result++
			}
		}
	}
	return result
}
