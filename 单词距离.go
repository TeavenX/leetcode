package main

import "math"

func main() {

}

func findClosest(words []string, word1 string, word2 string) int {
	result := math.MaxInt
	word1idx, word2idx := -1, -1
	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			word1idx = i
		}
		if words[i] == word2 {
			word2idx = i
		}
		if word1idx != -1 && word2idx != -1 {
			if result > abs(word2idx-word1idx) {
				result = abs(word1idx - word2idx)
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
