package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 6}
	nums = []int{38, 33, 57, 65, 13, 2, 86, 75, 4, 56}
	//fmt.Println(smallestDistancePair(nums, 3))
	fmt.Println(smallestDistancePair(nums, 26))
}

func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	if n == 2 {
		return abs(nums[0] - nums[1])
	}
	sort.Ints(nums)
	left, right := 0, nums[n-1]-nums[0]
	check := func(distance int) bool {
		count, left := 0, 0
		for right, num := range nums {
			for num-nums[left] > distance {
				left++
			}
			count += right - left
		}
		return count >= k
	}
	for left <= right {
		mid := (left + right) >> 1
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
