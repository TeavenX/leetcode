package main

import "fmt"

func main() {
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	cache := make(map[int]bool)
	for n > 1 {
		temp := 0
		for n > 0 {
			mul := n % 10
			temp += mul * mul
			n /= 10
		}
		if exist := cache[temp]; exist {
			return false
		}
		cache[temp] = true
		n = temp
	}
	return true
}
