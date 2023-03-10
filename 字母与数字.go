package main

import "fmt"

func main() {
	array := []string{"A", "B", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}
	fmt.Println(findLongestSubarray(array))
}

func findLongestSubarray(array []string) []string {
	preSum := make([]int, len(array)+1)
	cache := make(map[int]int)
	left, right, length := 0, 0, -1
	for i, item := range array {
		i++
		preSum[i] = preSum[i-1] + score(item)
		if preSum[i] == 0 {
			length = i
			right = i
		}
		if _, ok := cache[preSum[i]]; !ok {
			cache[preSum[i]] = i
		}
	}
	for r, s := range preSum {
		if l, ok := cache[s]; ok && r-l > length {
			length = r - l
			left = l
			right = r
		}
	}
	return array[left:right]
}

func score(v string) int {
	if v[0] >= '0' && v[0] <= '9' {
		return 1
	}
	return -1
}
