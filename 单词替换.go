package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	fmt.Println(replaceWords(dictionary, "the cattle was rattled by the battery"))
}

type dictTreeNode struct {
	cur  rune
	next []*dictTreeNode
}

func replaceWords(dictionary []string, sentence string) string {
	cache := make(map[string]bool)
	lencache := make(map[int]bool)
	for _, word := range dictionary {
		cache[word] = true
		lencache[len(word)] = true
	}
	lenlist := make([]int, 0)
	for k, _ := range lencache {
		lenlist = append(lenlist, k)
	}
	sort.Ints(lenlist)
	words := strings.Split(sentence, " ")
	for idx, word := range words {
		for _, l := range lenlist {
			if l <= len(word) {
				if exist := cache[word[:l]]; exist {
					words[idx] = word[:l]
					break
				}
			}
		}
	}
	return strings.Join(words, " ")
}
