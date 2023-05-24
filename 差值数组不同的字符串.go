package main

import (
	"bytes"
	"strconv"
)

func main() {

}

func oddString(words []string) string {
	cache := make(map[string][]int)
	n := len(words[0]) - 1
	for i, word := range words {
		b := bytes.Buffer{}
		for i := 1; i <= n; i++ {
			b.WriteString(strconv.Itoa(int(word[i] - word[i-1])))
			b.WriteByte(',')
		}
		s := b.String()
		cache[s] = append(cache[s], i)
	}
	for _, v := range cache {
		if len(v) == 1 {
			return words[v[0]]
		}
	}
	return ""
}

func oddStringV2(words []string) string {
	cache := make(map[string][]int)
	n := len(words[0]) - 1
	for i, word := range words {
		arr := make([]byte, n)
		for i := 1; i <= n; i++ {
			arr[i-1] = word[i] - word[i-1]
		}
		s := string(arr)
		cache[s] = append(cache[s], i)
	}
	for _, v := range cache {
		if len(v) == 1 {
			return words[v[0]]
		}
	}
	return ""
}
