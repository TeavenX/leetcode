package main

import "sort"

func main() {

}

func maxProduct(nums []int) int {
	sort.Ints(nums)
	a, b := (nums[0]-1)*(nums[1]-1), (nums[len(nums)-2]-1)*(nums[len(nums)-1]-1)
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
