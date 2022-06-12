package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	nums = []int{4, 5, 1, 8, 2}
	fmt.Println(productExceptSelf(nums))
	fmt.Println(productExceptSelfV2(nums))
	fmt.Println(productExceptSelfExplain(nums))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 2 {
		return []int{nums[1], nums[0]}
	}
	result := make([]int, n)
	left, right := make([]int, n), make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		result[i] = left[i] * right[i]
	}
	return result
}

func productExceptSelfV2(nums []int) []int {
	n := len(nums)
	if n == 2 {
		return []int{nums[1], nums[0]}
	}
	result := make([]int, n)
	left, right := 1, 1
	for i := 0; i < n; i++ {
		result[i] = 1
	}
	for i := 0; i < n; i++ {
		// 如果只是左边的循环，这里用=就可以
		// 但是现在把左边和右边放到了一个循环内
		// 所以result[i]可能包含了右边的值，所以需要用*=
		result[i] *= left
		left *= nums[i]

		result[n-1-i] *= right
		right *= nums[n-1-i]
	}
	return result
}

func productExceptSelfExplain(nums []int) []int {
	n := len(nums)
	if n == 2 {
		return []int{nums[1], nums[0]}
	}
	result := make([]int, n)
	left, right := 1, 1
	for i := 0; i < n; i++ {
		result[i] = 1
	}
	for i := 0; i < n; i++ {
		result[i] = left
		left *= nums[i]
	}
	for i := 0; i < n; i++ {
		result[n-1-i] *= right
		right *= nums[n-1-i]
	}
	return result
}
