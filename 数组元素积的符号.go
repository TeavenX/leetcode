package main

import "fmt"

func main() {
	nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	fmt.Println(arraySign(nums))
}

func arraySign(nums []int) int {
	m := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			m *= -1
		} else {
			m *= 1
		}
	}
	return m
}

func arraySignV2(nums []int) int {
	m := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		m ^= num
	}
	if m >= 0 {
		return 1
	}
	return -1
}
