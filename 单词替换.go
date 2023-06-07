package main

import (
	"bytes"
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

type Trie struct {
	Next  [26]*Trie
	isEnd bool
}

func (t *Trie) insert(s string) {
	node := t
	for _, b := range s {
		if node.Next[b-'a'] == nil {
			node.Next[b-'a'] = &Trie{}
		}
		node = node.Next[b-'a']
	}
	node.isEnd = true
}

func (t *Trie) search(s string) int {
	node := t
	for i, b := range s {
		if node.isEnd {
			return i
		}
		if node.Next[b-'a'] == nil {
			return len(s)
		}
		node = node.Next[b-'a']
	}
	return len(s)
}

func replaceWords(dictionary []string, sentence string) string {
	trie := Trie{}
	ss := strings.Split(sentence, " ")
	for _, dict := range dictionary {
		trie.insert(dict)
	}
	ans := bytes.Buffer{}
	for i, s := range ss {
		if i > 0 {
			ans.WriteByte(' ')
		}
		ans.WriteString(s[:trie.search(s)])
	}
	return ans.String()
}
