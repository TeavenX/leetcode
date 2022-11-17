package main

import "sort"

func main() {

}

const mod int = 1e9 + 7

func sumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	result := 0
	n := len(nums)
	cache := make([]int, n+1)
	cache[0] = 1
	for i := 1; i <= n; i++ {
		cache[i] = (cache[i-1] << 1) % mod
	}
	for i := 0; i < n; i++ {
		left := i
		right := n - 1 - i
		result = (result + (cache[left]-cache[right])*nums[i]) % mod
	}
	return (result + mod) % mod
}
