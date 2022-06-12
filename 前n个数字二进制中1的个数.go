package main

import "fmt"

func main() {
	//fmt.Println(countOne(7))
	fmt.Println(countBitsBest(50))
}

func countBitsBest(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = result[i>>1] + i&1
	}
	return result
}

func countBits(n int) []int {
	result := make([]int, n+1)
	result[0] = 0
	for i := 1; i <= n; i++ {
		result[i] = countOne(i)
	}
	return result
}

func countOne(n int) int {
	count := 0
	for n > 0 {
		count += 1
		n = n & (n - 1)
	}
	return count
}
