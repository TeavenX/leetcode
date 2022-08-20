package main

import "strings"

func main() {

}

func isPrefixOfWord(sentence string, searchWord string) int {
	sents := strings.Split(sentence, " ")
	for idx, sent := range sents {
		if len(sent) >= len(searchWord) {
			i := 0
			for i < len(searchWord) {
				if sent[i] != searchWord[i] {
					break
				}
				i++
			}
			if i == len(searchWord) {
				return idx
			}
		}
	}
	return -1
}
