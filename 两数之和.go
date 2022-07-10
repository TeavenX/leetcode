package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	nums = []int{3, 3}
	target := 6
	fmt.Println(twoSum20220613(nums, target))
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

func twoSum20220613(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}
	cache := make(map[int]int)
	for i, num := range nums {
		if idx, exist := cache[target-num]; exist && idx != i {
			return []int{i, idx}
		}
		cache[num] = i
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for idx, num := range nums {
		if i, ok := cache[target-num]; ok {
			if i != idx {
				return []int{i, idx}
			}
		}
		cache[num] = idx
	}
	return []int{}
}
