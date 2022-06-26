package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("dddccdbba"))
}

func firstUniqChar(s string) int {
	cache := make(map[byte]int)
	idx := -1
	for i := len(s) - 1; i >= 0; i-- {
		cache[s[i]]++
		if cache[s[i]] == 1 {
			idx = i
		} else if idx != -1 && s[idx] == s[i] {
			idx = -1
		}
	}
	return idx
}
