package main

import "fmt"

func main() {
	//nums := []int{2, 2, 1, 1, 5, 3, 3, 5}
	nums := []int{10, 2, 8, 9, 3, 8, 1, 5, 2, 3, 7, 6}
	fmt.Println(maxEqualFreq(nums))
}

func maxEqualFreq(nums []int) int {
	count := make(map[int]int)
	freq := make(map[int]int)
	maxFreq := 0
	result := 0
	for i, num := range nums {
		if count[num] > 0 {
			freq[count[num]]--
		}
		count[num]++
		maxFreq = max(maxFreq, count[num])
		freq[count[num]]++
		if maxFreq == 1 ||
			freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 ||
			freq[maxFreq]*maxFreq+1 == i+1 && freq[1] == 1 {
			result = max(result, i+1)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
