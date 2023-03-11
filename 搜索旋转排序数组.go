package main

import "fmt"

func main() {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// target := 0
	nums := []int{8, 9, 2, 3, 4}
	target := 9
	// nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
	// target := 0
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// target := 3
	// nums := []int{3, 1}
	// target := 1
	// nums := []int{1, 3}
	// target := 3
	// nums := []int{1}
	// target := 1
	// nums := []int{1}
	// target := 0
	// fmt.Println(search20220521(nums, target))
	fmt.Println(search20230312(nums, target))
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

func search20230311WA(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[0] {
			// mid落在左区间
			if target > nums[0] && target < nums[mid] {
				// target 也在左边，且target 比 mid要小，应该往左查找
				right = mid - 1
			} else {
				// target在右边，或者target比mid要大
				left = mid + 1
			}
		} else {
			// mid落在右区间
			if target < nums[0] && target > nums[mid] {
				// target 也在右边，且target比mid要大，应该往右找
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func search20230312(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	if nums[0] < nums[n-1] {
		left, right := 0, n-1
		for left <= right {
			mid := left + (right-left)>>1
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1
	} else {
		left := search(nums[:n>>1], target)
		if left != -1 {
			return left
		}
		right := search(nums[n>>1:], target)
		if right != -1 {
			return n>>1 + right
		}
	}
	return -1
}
