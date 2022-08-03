package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("abc" < "bcd")
}

func orderlyQueue(s string, k int) string {
	if k == 1 {
		result := s
		for i := 1; i < len(s); i++ {
			result = min(result, s[i:]+s[:i])
		}
		return result
	} else {
		str := []byte(s)
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		return string(str)
	}
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}
