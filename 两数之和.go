package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	nums = []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for idx, num := range nums {
		if idx2, exist := cache[target-num]; exist {
			return []int{idx, idx2}
		}
		cache[num] = idx
	}
	return []int{}
}

func twoSum20220504(nums []int, target int) []int {
	cache := make(map[int]int)
	for idx0, num := range nums {
		if idx, exist := cache[target-num]; exist {
			return []int{idx, idx0}
		}
		cache[num] = idx0
	}
	return []int{}
}
