package main

import (
	"sort"
	"strings"
)

func main() {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	removeSubfolders(folder)
}

type Trie struct {
	next  map[string]*Trie
	isEnd bool
}

func (t *Trie) insert(val string) {
	cur := t
	v := ""
	for i, b := range val[1:] {
		if b != '/' {
			v += string(b)
		}
		if b == '/' || i == len(val)-2 {
			if trie, ok := cur.next[v]; !ok {
				cur.next[v] = &Trie{
					next: make(map[string]*Trie),
				}
			} else if trie.isEnd {
				return
			}
			cur = cur.next[v]
			v = ""
		}
	}
	cur.isEnd = true
}

func (t *Trie) search(val string) bool {
	cur := t
	v := ""
	for i, b := range val[1:] {
		if b != '/' {
			v += string(b)
		}
		if b == '/' || i == len(val)-2 {
			trie := cur.next[v]
			if trie.isEnd && i < len(val)-2 {
				return false
			}
			cur = trie
			v = ""
		}
	}
	return true
}

func removeSubfolders(folder []string) []string {
	trie := &Trie{next: make(map[string]*Trie)}
	for _, f := range folder {
		trie.insert(f)
	}
	ans := make([]string, 0)
	for _, f := range folder {
		if trie.search(f) {
			ans = append(ans, f)
		}
	}
	return ans
}

func removeSubfoldersV2(folder []string) []string {
	sort.Strings(folder)
	ans := []string{folder[0]}
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			ans = append(ans, f)
		}
	}
	return ans
}
