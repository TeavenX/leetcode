package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 8, 7, -7, 5, 3, -7}
	fmt.Println(frequencySortV2(nums))
}

func frequencySort(nums []int) []int {
	cache := make(map[int][]int)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	for k, v := range freq {
		cache[v] = append(cache[v], k)
	}
	result := make([]int, 0, len(nums))
	for i := 1; i <= len(nums); i++ {
		if _, exist := cache[i]; exist {
			c := cache[i]
			sort.Slice(c, func(i, j int) bool {
				return c[j] < c[i]
			})
			for _, num := range c {
				for k := 0; k < freq[num]; k++ {
					result = append(result, num)
				}
			}
		}
	}
	return result
}

func frequencySortV2(nums []int) []int {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return cache[a] < cache[b] || (cache[a] == cache[b] && b < a)
	})
	return nums
}
