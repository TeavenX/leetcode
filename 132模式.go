package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-1, 3, 2, 0}
	fmt.Println(find132pattern(nums))
}

func find132pattern(nums []int) bool {
	j, k := make([]int, 0), math.MinInt
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		if len(j) == 0 {
			j = append(j, num)
		} else {
			for len(j) > 0 && j[len(j)-1] < num {
				k = j[len(j)-1]
				j = j[:len(j)-1]
			}
			j = append(j, num)
			if num < k {
				return true
			}
		}
	}
	return false
}
