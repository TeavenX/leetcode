package main

import "fmt"

func main() {
	nums1, nums2 := []int{4, 1, 2}, []int{1, 3, 4, 2}
	result := nextGreaterElement(nums1, nums2)
	fmt.Println(result)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	stack := make([]int, 0)
	cache := make(map[int]int)
	result := make([]int, n1)
	for i := n2 - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && stack[len(stack)-1] <= num {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			// 当前元素的下一个更大元素是stack中最后一个元素
			// 因为是倒序遍历，所以可以直接取（长板效应）
			cache[num] = stack[len(stack)-1]
		} else {
			cache[num] = -1
		}
		stack = append(stack, num)
	}
	for idx, num := range nums1 {
		result[idx] = cache[num]
	}
	return result
}
