package main

import (
	"math"
)

func main() {

}

func nextGreaterElement(n int) int {
	cache := make([]int, 0)
	for n > 0 {
		cache = append(cache, n%10)
		n /= 10
	}
	if len(cache) == 1 {
		return -1
	}
	for i := 1; i < len(cache); i++ {
		if cache[i] < cache[i-1] {
			for j := 0; j < i; j++ {
				if cache[j] > cache[i] {
					cache[i], cache[j] = cache[j], cache[i]
					reverse(cache, 0, i-1)
					goto next
				}
			}
		}
	}
	return -1
next:
	result := int64(0)
	for i := len(cache) - 1; i >= 0; i-- {
		result += int64(cache[i])
		result *= 10
	}
	result /= 10
	if result > math.MaxInt32 {
		return -1
	}
	return int(result)
}

func reverse(data []int, start, end int) {
	for start < end {
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}
}
