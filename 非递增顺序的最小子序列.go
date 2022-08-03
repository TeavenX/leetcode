package main

import "sort"

func main() {

}

func minSubsequence(nums []int) []int {
	sum := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for _, num := range nums {
		sum += num
	}
	half := (sum + 1) >> 1
	idx := 0
	for i, num := range nums {
		sum -= num
		idx = i
		if sum < half {
			break
		}
	}
	return nums[:idx]
}
