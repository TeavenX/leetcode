package main

import "sort"

func main() {

}

func specialArray(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	if n <= nums[0] {
		return n
	}
	for i := 1; i < n; i++ {
		if nums[n-i-1] < i && i <= nums[n-i] {
			return i
		}
	}
	return -1
}
