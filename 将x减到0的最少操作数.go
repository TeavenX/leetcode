package main

import "fmt"

func main() {
	nums := []int{3, 2, 20, 1, 1, 3}
	fmt.Println(minOperations(nums, 10))
}

func minOperations(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < x {
		return -1
	}
	target := sum - x
	n := len(nums)
	sum = 0
	left := 0
	ans := -1
	for right, num := range nums {
		sum += num
		for sum > target {
			sum -= nums[left]
			left++
		}
		if sum == target {
			t := n - (right - left + 1)
			if ans == -1 {
				ans = t
			} else {
				ans = min(ans, t)
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
