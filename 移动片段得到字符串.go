package main

import "strings"

func main() {

}

func canChange(start string, target string) bool {
	if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") {
		return false
	}
	n := len(target)
	j := 0
	for i := 0; i < len(start); i++ {
		if start[i] != '_' {
			for j < n && target[j] == '_' {
				j++
			}
			if j >= n {
				return false
			}
			if start[i] == 'L' && i < j {
				return false
			}
			if start[i] == 'R' && i > j {
				return false
			}
			j++
		}
	}
	return true
}
