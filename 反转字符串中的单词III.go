package main

import "fmt"

func main() {
	fmt.Println(reverseWordsV2("God Ding"))
}

func reverseWords(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		temp := ""
		for i < len(s) && s[i] != ' ' {
			temp = string(s[i]) + temp
			i++
		}
		result += temp
		if i < len(s) {
			result += " "
		}
	}
	return result
}

func reverseWordsV2(s string) string {
	sentence := []byte(s)
	start := 0
	reverse := func(start, end int) {
		for i := 0; i <= (end-start)>>1; i++ {
			sentence[start+i], sentence[end-i] = sentence[end-i], sentence[start+i]
		}
	}
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			reverse(start, i-1)
			start = i + 1
		}
	}
	reverse(start, len(sentence)-1)
	return string(sentence)
}
