package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 6
	//nums = []int{3, 1}
	//target = 1
	fmt.Println(search20220521(nums, target))
}

func search(nums []int, target int) int {
	result := -1
	n := len(nums)
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return result
		}
	}
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			// left 和 mid在同一侧（同左或者同右）
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// left 在左侧，mid在右侧
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return result
}

func search20220521(nums []int, target int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else {
			if nums[mid] >= nums[left] {
				// mid 和 left在同一侧
				if nums[mid] > target && nums[left] <= target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				if nums[mid] < target && nums[right] >= target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}
	return -1
}
