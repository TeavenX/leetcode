package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumSwap(2736))
}

func maximumSwap(num int) int {
	if num <= 11 {
		return num
	}
	arr := make([]int, 0)
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	cache := append([]int{}, arr...)
	sort.Ints(cache)
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		if arr[i] != cache[i] {
			for j := range arr {
				if arr[j] == cache[i] {
					arr[i], arr[j] = arr[j], arr[i]
					goto next
				}
			}
		}
	}
next:
	result := 0
	for i := n - 1; i >= 0; i-- {
		result *= 10
		result += arr[i]
	}
	return result
}
