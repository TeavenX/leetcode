package main

import "fmt"

func main() {
	a := []int{1, 1, 1, 2, 2, 2, 3}
	fmt.Println(singleNumber(a))
}

func singleNumber(nums []int) int {
	var temp [32]int32
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			temp[i] += int32(num) & 1
			num >>= 1
		}
	}
	result := int32(0)
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= temp[31-i] % 3
	}
	return int(result)
}

func singleNumberTwiceTimes(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	result := 0
	for i := 0; i < n; i++ {
		result = result ^ nums[i]
	}
	return result
}

func singleNumber20220524(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func singleNumber20220524V2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	result := nums[0]
	for i := 1; i < n; i++ {
		result ^= nums[i]
	}
	return result
}
