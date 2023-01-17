package main

import "fmt"

func main() {
	nums := []int{13, 10, 35, 24, 76}
	fmt.Println(countNicePairs(nums))
}

const mod int = 1e9 + 7

func countNicePairs(nums []int) int {
	cache := make(map[int]int)
	ans := 0
	for _, num := range nums {
		r, t := 0, num
		for t > 0 {
			r *= 10
			r += t % 10
			t /= 10
		}
		fmt.Println(num, r, num-r)
		if _, ok := cache[num-r]; ok {
			ans = (ans + cache[num-r]) % mod
			fmt.Println("ans++")
		}
		cache[num-r]++
	}
	fmt.Println(cache)
	return ans
}
