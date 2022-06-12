package main

import (
	"fmt"
	"strings"
)

func main() {

	words := []string{"abc", "test", "qqq", "mcxfjdas", "ab"}
	//fmt.Println(maxProduct(words))
	//fmt.Println(maxProductV2(words))
	fmt.Println(maxProductV3(words))

}

func maxProduct(words []string) int {
	result := 0
	for i := 0; i < len(words); i++ {
		word1 := words[i]
		for j := i + 1; j < len(words); j++ {
			word2 := words[j]
			if !hasSameCharV3(words[j], word1) {
				lenC := len(word1) * len(word2)
				if lenC > result {
					result = lenC
				}
			}
		}
	}
	return result
}

func maxProductV2(words []string) int {
	result := 0
	numWords := len(words)
	masks := make([]int, numWords)
	for i := 0; i < numWords; i++ {
		bitMask := 0
		for _, c := range words[i] {
			bitMask |= (1 << (c - 'a'))
		}
		masks[i] = bitMask
	}
	for i := 0; i < numWords; i++ {
		for j := i + 1; j < numWords; j++ {
			if masks[i]&masks[j] == 0 {
				lenC := len(words[i]) * len(words[j])
				if lenC > result {
					result = lenC
				}
			}
		}
	}
	return result
}
func maxProductV3(words []string) int {
	result := 0
	numWords := len(words)
	masks := map[int]int{}
	for i := 0; i < numWords; i++ {
		bitMask := 0
		for _, c := range words[i] {
			bitMask |= 1 << (c - 'a')
		}
		if len(words[i]) > masks[bitMask] {
			masks[bitMask] = len(words[i])
		}
	}
	for i := range masks {
		for j := range masks {
			if i&j == 0 {
				lenC := masks[i] * masks[j]
				if lenC > result {
					result = lenC
				}
			}
		}
	}
	return result
}

func hasSameChar(word1, word2 string) bool {
	for _, i := range word1 {
		if strings.Index(word2, string(i)) != -1 {
			return true
		}
	}
	return false
}

func hasSameCharV2(word1, word2 string) bool {
	var temp [26]int
	for _, i := range word1 {
		temp[i-'a'] = 1
	}
	for _, i := range word2 {
		if temp[i-'a'] == 1 {
			return true
		}
	}
	return false
}

func hasSameCharV3(word1, word2 string) bool {
	bitMask1, bitMask2 := 0, 0
	for _, i := range word1 {
		bitMask1 |= 1 << (i - 'a')
	}
	for _, i := range word2 {
		bitMask2 |= 1 << (i - 'a')
	}
	if bitMask1&bitMask2 == 0 {
		return false
	}
	return true
}
