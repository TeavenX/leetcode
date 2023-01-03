package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxValue(6, 3, 15))
}

func calc(top, length int) int {
	if length == 0 {
		return 0
	}
	if length <= top {
		// 可以分到不止1个数
		return (top - length + 1 + top) * length / 2
	}
	return (1+top)*top/2 + length - top
}

func maxValue(n int, index int, maxSum int) int {
	left, right := index, n-index-1
	return sort.Search(maxSum, func(mid int) bool {
		return mid+calc(mid, left)+calc(mid, right) >= maxSum
	})
}
