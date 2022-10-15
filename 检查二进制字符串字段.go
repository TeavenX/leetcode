package main

import "strings"

func main() {

}

func checkOnesSegment(s string) bool {
	exist := false
	for i, n := range s {
		if n == '1' {
			if exist && i > 0 && s[i-1] == '0' {
				return false
			}
			exist = true
		}
	}
	return true
}

func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}
