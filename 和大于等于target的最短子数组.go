package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	nums = []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}
	target = 213
	result := minSubArrayLen(target, nums)
	fmt.Println(result)
}

func minSubArrayLen(target int, nums []int) int {
	left, sum, result := 0, 0, len(nums)+1
	for right, num := range nums {
		sum += num
		for sum >= target {
			newSize := right - left + 1
			if newSize < result {
				result = newSize
			}
			sum -= nums[left]
			left++
		}
	}
	if result > len(nums) {
		result = 0
	}
	return result
}
