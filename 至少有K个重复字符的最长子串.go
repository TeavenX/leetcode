package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(longestSubstring("aaabb", 3))
	fmt.Println(longestSubstring("bbaaacbd", 3))
}

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	cache := [26]int{}
	for _, b := range s {
		cache[b-'a']++
	}
	for b, c := range cache {
		if c > 0 && c < k {
			ans := 0
			for _, t := range strings.Split(s, string(byte(b+'a'))) {
				ans = max(ans, longestSubstring(t, k))
			}
			return ans
		}
	}
	bytes.NewBuffer(nil).WriteString()
	return len(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
