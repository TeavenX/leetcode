package main

import "sort"

func main() {

}

const mod int = 1e9 + 7

func sumDistance(nums []int, s string, d int) int {
	for i := range nums {
		if s[i] == 'L' {
			nums[i] -= d
		} else {
			nums[i] += d
		}
	}
	sort.Ints(nums)
	ans := 0
	sum := 0
	for i, num := range nums {
		ans = (ans + i*num - sum) % mod
		sum = sum + num
	}
	return ans
}
