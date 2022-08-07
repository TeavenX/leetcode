package main

import (
	"sort"
	"strings"
)

func main() {

}

func stringMatching(words []string) []string {
	n := len(words)
	if n == 1 {
		return []string{}
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	result := make([]string, 0)
next:
	for i := 0; i < n-1; i++ {
		sub := words[i]
		for j := i + 1; j < n; j++ {
			if strings.Contains(words[j], sub) {
				result = append(result, sub)
				continue next
			}
		}
	}
	return result
}

func match(a, b string) bool {
	lenA, lenB := len(a), len(b)
	ia, ib := 0, 0
	for ib < lenB-lenA {
		for b[ib] == a[ia] && ia < lenA {
			ib++
			ia++
		}
		if ia == lenA {
			return true
		}
	}
	return false
}
