package main

import "sort"

func main() {

}

func singleNumber(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[n-1] != nums[n-2] {
		return nums[n-1]
	}
	for i, num := range nums {
		if i == 0 || i == n-1 {
			continue
		}
		if num != nums[i-1] && num != nums[i+1] {
			return num
		}
	}
	return -1
}

func singleNumberV2(nums []int) int {
	// 卡诺图
	// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/comments/931045
	a, b := 0, 0
	for _, c := range nums {
		t := a
		a = b&c | a&(^c)
		b = b&(^c) | (^t)&(^b)&c
	}
	return b
}
