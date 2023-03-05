package main

import "strings"

func main() {

}

func minimumDeletions(s string) int {
	del := strings.Count(s, "a")
	ans := del
	for _, b := range s {
		if b == 'a' {
			del--
		} else {
			del++
		}
		// del += int(b-'a') * 2 - 1
		if del < ans {
			ans = del
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
