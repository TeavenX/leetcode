package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) == 1 {
		return true
	}
	cache := make(map[byte]byte)
	cacheMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := cache[s[i]]; !ok {
			cache[s[i]] = t[i]
			if m, ok := cacheMap[t[i]]; ok {
				if m != s[i] {
					return false
				}
			}
			cacheMap[t[i]] = s[i]
		} else {
			if v != t[i] {
				return false
			}
		}
	}
	return true
}
