package main

import "fmt"

func main() {
	piles := []int{312884470}
	h := 968709470
	fmt.Println(minEatingSpeed(piles, h))
}

func minEatingSpeed(piles []int, h int) int {
	var left int = 0
	var right int = 1e9
	check := func(k int) bool {
		result := 0
		for _, pile := range piles {
			result += (pile + k - 1) / k
		}
		return result <= h
	}
	for left < right {
		mid := (left + right) >> 1
		if mid == 0 {
			return 1
		}
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
