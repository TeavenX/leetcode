package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	//nums = []int{1, 2, 3, 4, 5}
	//target = 11
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	left, sum, result := 0, 0, 0
	for right, num := range nums {
		if num >= target {
			return 1
		} else {
			sum += num
			for sum >= target {
				cur := right - left + 1
				if result == 0 || cur < result {
					result = cur
				}
				sum -= nums[left]
				left++
			}
		}
	}
	return result
}
