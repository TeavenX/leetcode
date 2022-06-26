package main

import "fmt"

func main() {
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring("barfoothefoobarman", words))
}

func findSubstring(s string, words []string) []int {
	cache := make(map[string]int)
	result := make([]int, 0)
	wlen := len(words[0])
	left := 0
	for _, word := range words {
		cache[word]++
	}
next:
	for right := wlen*len(words) - 1; right < len(s); right++ {
		match := make(map[string]int)
		for i := left; i <= right; i += wlen {
			str := s[i : i+wlen]
			if count, exist := cache[str]; !exist {
				left++
				continue next
			} else {
				mc, _ := match[str]
				if mc < count {
					match[str]++
				} else {
					left++
					continue next
				}
			}
		}
		result = append(result, left)
		left++
	}
	return result
}
