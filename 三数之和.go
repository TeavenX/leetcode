package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	nums = []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				right--
				left++
				for nums[right] == nums[right+1] && right > left {
					right--
				}
				for nums[left] == nums[left-1] && right > left {
					left++
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func threeSum20220427(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return make([][]int, 0)
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return result
}

func threeSum20220504(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	result := make([][]int, 0)
	if n < 3 {
		return result
	}
	for idx, num := range nums {
		if idx > 0 && num == nums[idx-1] {
			continue
		}
		target := -num
		left, right := idx+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{num, nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
