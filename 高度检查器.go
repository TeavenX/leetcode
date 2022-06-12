package main

import "sort"

func main() {

}

func heightChecker(heights []int) int {
	n := len(heights)
	if n == 1 {
		return 0
	}
	sortH := make([]int, n)
	copy(sortH, heights)
	sort.Ints(sortH)
	result := 0
	for i, num := range sortH {
		if num != heights[i] {
			result++
		}
	}
	return result
}

func heightCheckerV2(heights []int) int {
	n := len(heights)
	if n == 1 {
		return 0
	}
	var cache [101]int
	for _, num := range heights {
		cache[num]++
	}
	result := 0
	j := 0
	for i, num := range cache {
		for ; num > 0; num-- {
			if heights[j] != i {
				result++
			}
			j++
		}
	}
	return result
}
