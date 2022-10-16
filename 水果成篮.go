package main

import "fmt"

func main() {
	//fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fruits := []int{1, 0, 1, 4, 1, 4, 1, 2, 3}
	fmt.Println(totalFruit(fruits))
}

func totalFruit(fruits []int) int {
	left := 0
	max := 0
	cache := make(map[int]int)
	for right, fruit := range fruits {
		if cache[fruit] == 0 {
			for len(cache) == 2 {
				cache[fruits[left]]--
				if cache[fruits[left]] == 0 {
					delete(cache, fruits[left])
				}
				left++
			}
		}
		cache[fruit]++
		if max < right-left+1 {
			max = right - left + 1
		}
	}
	return max
}
