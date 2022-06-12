package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := make([]int, 0)
	target := 0
	nums = []int{2, 2, 2, 2, 2, 2}
	nums = []int{1, 0, -1, 0, -2, 2}
	target = 8
	target = 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		targetL1 := target - nums[i]
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			targetL2 := targetL1 - nums[j]
			left, right := j+1, n-1
			for left < right {
				sum := nums[left] + nums[right]
				if sum == targetL2 {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for nums[left] == nums[left-1] && left < right {
						left++
					}
					for nums[right] == nums[right+1] && left < right {
						right--
					}
				} else if sum > targetL2 {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}

func fourSum20220504(nums []int, target int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	if n < 4 {
		return result
	}
	sort.Ints(nums)
	for idx, num := range nums {
		if idx > 0 && num == nums[idx-1] {
			continue
		}
		for i := idx + 1; i < n; i++ {
			if i > idx+1 && nums[i] == nums[i-1] {
				continue
			}
			newt := target - num - nums[i]
			left, right := i+1, n-1
			for left < right {
				sum := nums[left] + nums[right]
				if sum == newt {
					result = append(result, []int{num, nums[i], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum > newt {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}
