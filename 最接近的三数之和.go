package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{0, 2, 1, -3}
	target := 1
	//nums = []int{1, 1, -1, -1, 3}
	//target = -1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	min := math.MaxInt
	result := -1
	for i := 0; i <= n-3; i++ {
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			dist := abs(target - sum)
			if dist < min {
				min = dist
				result = sum
			} else {
				if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
