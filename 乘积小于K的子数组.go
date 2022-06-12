package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	k := 1
	nums = []int{57, 44, 92, 28, 66, 60, 37, 33, 52, 38, 29, 76, 8, 75, 22}
	k = 18
	//nums = []int{100, 2, 3, 4, 100, 5, 6, 7, 100}
	//k = 100
	result := numSubarrayProductLessThanKV2(nums, k)
	fmt.Println(result)
}

func numSubarrayProductLessThanKV1(nums []int, k int) int {
	n := len(nums)
	result := 0
	for left, num := range nums {
		sum := num
		right := left + 1
		if num < k {
			result++
		}
		for right < n && sum < k {
			sum = sum * nums[right]
			if sum < k {
				result++
			} else {
				break
			}
			right++
		}
	}
	return result
}

func numSubarrayProductLessThanKV2(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	result, left, sum := 0, 0, 1
	for right, num := range nums {
		sum *= num
		for sum >= k {
			sum /= nums[left]
			left++
		}
		result += right - left + 1
	}
	return result
}
