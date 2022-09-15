package main

import "fmt"

func main() {
	fmt.Println(uniqueLetterString("LEETCODE"))
}

func uniqueLetterString(s string) int {
	left, right := make([]int, len(s)), make([]int, len(s))
	pre := [26]int{}
	for i := 0; i < 26; i++ {
		pre[i] = -1
	}
	for i, str := range s {
		s := str - 'A'
		left[i] = pre[s]
		pre[s] = i
	}
	for i := 0; i < 26; i++ {
		pre[i] = len(s)
	}
	for i := len(s) - 1; i >= 0; i-- {
		str := s[i] - 'A'
		right[i] = pre[str]
		pre[str] = i
	}
	result := 0
	for i := 0; i < len(s); i++ {
		result += (right[i] - i) * (i - left[i])
	}
	return result
}
