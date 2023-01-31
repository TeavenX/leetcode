package main

import "bytes"

func main() {

}

func decodeMessage(key string, message string) string {
	dict := make(map[rune]byte)
	idx := 0
	for _, b := range key {
		if b == ' ' {
			continue
		}
		if _, exist := dict[b]; !exist {
			dict[b] = 'a' + byte(idx)
			idx++
		}
		if idx == 26 {
			break
		}
	}
	dict[' '] = ' '
	ans := bytes.NewBufferString("")
	for _, b := range message {
		ans.WriteByte(dict[b])
	}
	return ans.String()
}
