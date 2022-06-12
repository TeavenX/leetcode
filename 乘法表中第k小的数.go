package main

import "fmt"

func main() {
	fmt.Println(findKthNumber(3, 3, 5))
}

func findKthNumber(m int, n int, k int) int {
	left, right := 1, m*n
	for left < right {
		mid := left + (right-left)>>1
		if countLineLessK(m, n, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func countLineLessK(m, n, mid int) int {
	result := 0
	for i := 1; i <= m; i++ {
		result += min(mid/i, n)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
