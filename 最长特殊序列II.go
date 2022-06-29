package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"aabbcc", "aabbcc", "cb"}
	strs = []string{"aabbcc", "aabbcc", "bc", "bcc", "aabbccc"}
	//strs = []string{"a", "b", "c", "d", "e", "f", "a", "b", "c", "d", "e", "f"}
	fmt.Println(findLUSlength(strs))
}

func findLUSlengthError(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		if len(strs[i]) > len(strs[j]) {
			return true
		} else if len(strs[i]) == len(strs[j]) {
			return strs[i][0]-strs[j][0] > 0
		}
		return false
	})
	preMatch := false
	for i := 0; i < len(strs)-1; i++ {
		str := strs[i]
		pattern := strs[i+1]
		p1, p2 := 0, 0
		for p1 < len(str) && p2 < len(pattern) {
			if str[p1] == pattern[p2] {
				p1++
				p2++
			} else {
				p1++
			}
		}
		if p2 < len(pattern) {
			return len(pattern)
		} else if p2 == len(pattern) && !preMatch && p1 < len(str) {
			return len(str)
		}
		preMatch = true
	}
	return -1
}

func findLUSlength(strs []string) int {
	result := -1
next:
	for i, str := range strs {
		for j, pattern := range strs {
			if i != j && isSubStr(str, pattern) {
				continue next
			}
		}
		if len(str) > result {
			result = len(str)
		}
	}
	return result
}

func isSubStr(a, b string) bool {
	if a == b {
		return true
	}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == len(a)
}
