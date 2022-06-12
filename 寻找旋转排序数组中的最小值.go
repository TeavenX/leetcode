package main

import "fmt"

func main() {
	nums := []int{3, 1, 2}
	fmt.Println(findMinV2(nums))
}
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	min := nums[left]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[left] {
			// mid和left在同一侧，同在左边或者右边
			if nums[mid] > nums[right] {
				// mid和left在左边，right在右边
				left = mid + 1
			} else {
				// left mid right在同一侧，左边或者右边
				return nums[left]
			}
		} else {
			// left 在左边，mid和right在右边
			right = mid
		}
	}
	return min
}

func findMinV2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[right]
}
