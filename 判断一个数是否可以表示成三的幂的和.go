package main

import "fmt"

func main() {
	r := 1
	i := 0
	for r < 1e7 {
		i++
		r *= 3
	}
	fmt.Println(i)
}

func checkPowersOfThree(n int) bool {
	cache := [15]int{}
	for i, t := 0, 1; i < 15; i++ {
		cache[i] = t
		t *= 3
	}
	i := 14
	for n > 0 && i >= 0 {
		if cache[i] <= n {
			n -= cache[i]
		}
		i--
	}
	return n == 0
}

func checkPowersOfThreeV2(n int) bool {
	// 3进制
	for ; n > 0; n /= 3 {
		if n%3 == 2 {
			return false
		}
	}
	return true
}
