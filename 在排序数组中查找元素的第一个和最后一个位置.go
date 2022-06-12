package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	n := len(nums)
	if n < 1 {
		return result
	}
	result[0] = searchIdx(nums, target, "left")
	result[1] = searchIdx(nums, target, "right")
	return result
}

func searchIdx(nums []int, target int, searchType string) int {
	n := len(nums)
	left, right := 0, n-1
	temp := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			temp = mid
			switch searchType {
			case "left":
				right = mid - 1
			case "right":
				left = mid + 1
			default:
				break
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return temp
}
