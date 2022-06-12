package main

import "fmt"

func main() {
	words := []string{"apap", "app"}
	order := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
}

func isAlienSorted(words []string, order string) bool {
	dict := make(map[byte]int, 26)
	for i := 0; i < 26; i++ {
		dict[order[i]] = i
	}
next:
	for i := 0; i < len(words)-1; i++ {
		pre := words[i]
		cur := words[i+1]
		lens := len(pre)
		if len(cur) < lens {
			lens = len(cur)
		}
		for j := 0; j < lens; j++ {
			if dict[pre[j]] > dict[cur[j]] {
				return false
			}
			if dict[pre[j]] < dict[cur[j]] {
				continue next
			}
		}
		if len(pre) > len(cur) {
			return false
		}
	}
	return true
}

func isAlienSortedV2(words []string, order string) bool {
	index := [26]int{}
	for i, c := range order {
		index[c-'a'] = i
	}
next:
	for i := 1; i < len(words); i++ {
		for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ {
			pre, cur := index[words[i-1][j]-'a'], index[words[i][j]-'a']
			if pre > cur {
				return false
			}
			if pre < cur {
				continue next
			}
		}
		if len(words[i-1]) > len(words[i]) {
			return false
		}
	}
	return true
}
