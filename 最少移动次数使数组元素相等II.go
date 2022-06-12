package main

import "sort"

func main() {

}

func minMoves2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	sort.Ints(nums)
	result := 0
	mid := nums[n/2]
	for _, num := range nums {
		result += abs(num - mid)
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}
