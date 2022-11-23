package main

import "fmt"

func main() {
	nums := []int{2, 9, 2, 5, 6}
	fmt.Println(numSubarrayBoundedMax(nums, 2, 8))
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	l := 0
	result := 0
	pre := 0
	for r, num := range nums {
		if num > right {
			l = r + 1
			result += pre
			pre = 0
			continue
		}
		if num >= left && num <= right {
			pre += r - l + 1
			// l = r + 1
		} else if pre != 0 {
			pre += pre
		}
	}
	if pre != 0 {
		result += pre
	}
	return result
}
