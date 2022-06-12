package main

import (
	"fmt"
)

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	fmt.Println(findAndReplacePatternV2(words, "abb"))
	words = []string{"ef", "fq", "ao", "at", "lx"}
	fmt.Println(findAndReplacePatternV2(words, "ya"))
}

func findAndReplacePattern(words []string, pattern string) []string {
	if len(words[0]) == 1 {
		return words
	}
	result := make([]string, 0)
next:
	for _, word := range words {
		cachelet := make(map[byte]byte)
		cachepat := make(map[byte]byte)
		for i := 0; i < len(word); i++ {
			letter := word[i]
			pat := pattern[i]
			if tlet, exist := cachepat[pat]; exist {
				if letter != tlet {
					continue next
				}
			} else {
				cachepat[pat] = letter
				if tpat, exist := cachelet[letter]; exist {
					if tpat != pat {
						continue next
					}
				}
				cachelet[letter] = pat
			}
		}
		result = append(result, word)
	}
	return result
}

func findAndReplacePatternV2(words []string, pattern string) []string {
	if len(words[0]) == 1 {
		return words
	}
	result := make([]string, 0)
	for _, word := range words {
		if match(word, pattern) && match(pattern, word) {
			result = append(result, word)
		}
	}
	return result
}

func match(word, pattern string) bool {
	cache := make(map[byte]byte)
	for i := 0; i < len(word); i++ {
		if pat, exist := cache[word[i]]; exist {
			if pat != pattern[i] {
				return false
			}
		} else {
			cache[word[i]] = pattern[i]
		}
	}
	return true
}
