package main

import "math"

func main() {
	nums := []int{2, 7, 10, 9, 8, 9}
	movesToMakeZigzag(nums)
}

func movesToMakeZigzag(nums []int) int {
	s := [2]int{}
	for i, num := range nums {
		left, right := math.MaxInt, math.MaxInt
		if i > 0 {
			left = nums[i-1]
		}
		if i < len(nums)-1 {
			right = nums[i+1]
		}
		s[i%2] += max(0, num-min(left, right)+1)
	}
	return min(s[0], s[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
