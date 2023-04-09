package main

import "sort"

func main() {

}

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	n := len(stones)
	emp1 := stones[n-1] - stones[1] - 1 - (n - 3)
	emp2 := stones[n-2] - stones[0] - 1 - (n - 3)
	maxMove := max(emp1, emp2)
	if emp1 == 0 || emp2 == 0 {
		return []int{min(2, maxMove), maxMove}
	}
	maxCount, left := 0, 0
	for right, x := range stones {
		for stones[left] <= x-n {
			left++
		}
		maxCount = max(maxCount, right-left+1)
	}
	return []int{n - maxCount, maxMove}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
