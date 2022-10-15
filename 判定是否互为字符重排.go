package main

import "sort"

func main() {

}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	cache := make(map[rune]int)
	for _, s := range s1 {
		cache[s]++
	}
	for _, s := range s2 {
		cache[s]--
		if cache[s] < 0 {
			return false
		}
	}
	return true
}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	b1, b2 := []byte(s1), []byte(s2)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}
