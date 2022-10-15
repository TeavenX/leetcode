package main

import "fmt"

func main() {
	fmt.Println(isFlipedString("waterbottle", "erbottlewat"))
}

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	n := len(s1)
	i, j := 0, 0
	for i < len(s1) {
		start := i
		for s1[i%n] == s2[j%n] {
			i++
			j++
			if i%n == start {
				return true
			}
		}
		i++
	}
	return false
}
