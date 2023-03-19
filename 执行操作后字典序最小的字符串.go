package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findLexSmallestString("5525", 9, 2))
	fmt.Println(findLexSmallestString("43987654", 7, 3))
}

func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	ans := s
	for i := 0; i < n; i++ {
		s = s[n-b:] + s[:n-b]
		cache := []byte(s)
		for j := 0; j < 10; j++ {
			for i, by := range cache {
				if i&1 == 1 {
					cache[i] = (by-'0'+byte(a))%10 + '0'
				}
			}
			if s = string(cache); s < ans {
				ans = s
			}
			if b&1 == 1 {
				for j := 0; j < 10; j++ {
					for i, by := range cache {
						if i&1 == 0 {
							cache[i] = (by-'0'+byte(a))%10 + '0'
						}
					}
					if s = string(cache); s < ans {
						ans = s
					}
				}
			}
		}
	}
	return ans
}
