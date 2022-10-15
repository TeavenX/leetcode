package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := []int{1}
	nums := []int{2, 3}
	fmt.Println(missingTwo(nums))
}

func missingTwoError(nums []int) []int {
	// 要求O（n）的时间复杂度，题目给的nums是乱序的，
	n := len(nums)
	dist := 0
	result := make([]int, 0, 2)
	for i := 1; i <= n+2; i++ {
		if i-1-dist >= n || nums[i-1-dist] != i {
			dist++
			result = append(result, i)
		}
	}
	return result
}

func missingTwo(nums []int) []int {
	n := len(nums) + 2
	sum := 0
	sqrtSum := 0
	for _, num := range nums {
		sum += num
		sqrtSum += num * num
	}
	ss := 0
	for i := 1; i <= n; i++ {
		ss += i * i
	}
	k := (1+n)*n>>1 - sum
	q := ss - sqrtSum
	t := int(math.Sqrt(float64(2*q - k*k)))
	return []int{(k + t) >> 1, (k - t) >> 1}
}

func missingTwoV2(nums []int) []int {
	n := len(nums) + 2
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	for i := 1; i < n; i++ {
		xorSum ^= i
	}

}
