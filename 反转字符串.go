package main

import "fmt"

func main() {
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(a)
	reverseString(a)
	fmt.Println(a)
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
