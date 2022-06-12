package main

import (
	"math"
	"sort"
)

func main() {

}

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	if n == 1 {
		return []int{-1}
	}
	result := make([]int, n)
	for i, itvl := range intervals {
		pre := math.MaxInt
		for idx, val := range intervals {
			if itvl[1] <= val[0] && pre > val[0] {
				pre = val[0]
				result[i] = idx
			}
		}
		if pre == math.MaxInt {
			result[i] = -1
		}
	}
	return result
}

func findRightIntervalV2(intervals [][]int) []int {
	n := len(intervals)
	if n == 1 {
		return []int{-1}
	}
	cache := make(map[int]int)
	prefix := make([]int, n)
	result := make([]int, n)
	for i, val := range intervals {
		cache[val[0]] = i
		prefix[i] = val[0]
	}
	sort.Ints(prefix)
	for i, val := range intervals {
		result[i] = -1
		left, right := 0, n-1
		for left <= right {
			mid := left + (right-left)>>1
			if prefix[mid] >= val[1] {
				result[i] = cache[prefix[mid]]
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return result
}
