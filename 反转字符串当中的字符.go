package main

import "fmt"

func main() {
	a := "Let's take LeetCode contest"
	ret := reverseWords(a)
	fmt.Println(ret)
}

func reverseWords(s string) string {
	result := ""
	lenS := len(s)
	right := lenS - 1
	temp := ""
	for right >= 0 {
		if s[right] == ' ' {
			if result != "" {
				result = " " + result
			}
			result = temp + result
			temp = ""
			right--
		} else {
			temp += string(s[right])
			right--
		}
	}
	if result != "" {
		result = " " + result
	}
	result = temp + result
	return result
}
