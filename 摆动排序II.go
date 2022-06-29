package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	nums = []int{1, 2, 3}
	wiggleSort(nums)
	fmt.Println(nums)
}

func wiggleSort(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	sort.Ints(nums)
	newNums := make([]int, n)
	mid := (n + 1) >> 1
	left, right := mid-1, n-1
	flag := true
	for i := 0; i < n; i++ {
		if flag && left >= 0 {
			newNums[i] = nums[left]
			left--
		} else if right >= mid {
			newNums[i] = nums[right]
			right--
		}
		flag = !flag
	}
	for i := 0; i < n; i++ {
		nums[i] = newNums[i]
	}
}
