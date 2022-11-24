package main

import "fmt"

func main() {
	// words := []string{"hello", "hi", "helo"}
	// fmt.Println(expressiveWords("heeellooo", words))
	// words := []string{"tl", "tll", "ttll", "ttl"}
	// fmt.Println(expressiveWords("tttttllll", words))
	words := []string{"aaaa"}
	fmt.Println(expressiveWords("aaa", words))
}

func expressiveWords(s string, words []string) int {
	result := 0
	for _, word := range words {
		if expand(s, word) {
			result++
		}
	}
	return result
}

func expand(s1, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	p1, p2 := 0, 0
	for p1 < l1 && p2 < l2 {
		if s1[p1] != s2[p2] {
			return false
		}
		c := s1[p1]
		count1, count2 := 0, 0
		for p1 < l1 && s1[p1] == c {
			p1++
			count1++
		}
		for p2 < l2 && s2[p2] == c {
			p2++
			count2++
		}
		if count1 < count2 || count1 > count2 && count1 < 3 {
			return false
		}
	}
	return p1 == l1 && p2 == l2
}
