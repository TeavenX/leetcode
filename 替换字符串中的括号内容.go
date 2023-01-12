package main

import "bytes"

func main() {

}

type Trie struct {
	next  [26]*Trie
	isEnd bool
	value string
}

func (t *Trie) Search(s string) string {
	node := t
	for _, b := range s {
		if node.next[b-'a'] == nil {
			return "?"
		}
		node = node.next[b-'a']
	}
	if !node.isEnd {
		return "?"
	}
	return node.value
}

func (t *Trie) Insert(k, v string) {
	node := t
	for _, b := range k {
		if node.next[b-'a'] == nil {
			node.next[b-'a'] = &Trie{}
		}
		node = node.next[b-'a']
	}
	node.isEnd = true
	node.value = v
}

func evaluate(s string, knowledge [][]string) string {
	trie := Trie{}
	for _, k := range knowledge {
		trie.Insert(k[0], k[1])
	}
	ans := bytes.NewBufferString("")
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			j := i
			for ; s[j] != ')'; j++ {
			}
			ans.WriteString(trie.Search(s[i+1 : j]))
			i = j
		} else {
			ans.WriteByte(s[i])
		}
	}
	return ans.String()
}
