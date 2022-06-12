package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	cache := make(map[int]bool)
	for _, num := range nums {
		if exist := cache[num]; exist {
			return true
		}
		cache[num] = true
	}
	return false
}
