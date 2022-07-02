package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}

func longestPalindrome(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}
	cache := make(map[rune]int)
	for _, str := range s {
		cache[str]++
	}
	result := 0
	for _, v := range cache {
		result += (v >> 1) << 1
	}
	if n-result > 0 {
		result++
	}
	return result
}
