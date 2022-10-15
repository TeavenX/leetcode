package main

import "strings"

func main() {

}

func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}
	if strings.Replace(start, "X", "", -1) != strings.Replace(end, "X", "", -1) {
		return false
	}
	n := len(start)
	i, j := 0, 0
	for i < n || j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}
		if i < n {
			if start[i] == 'L' && i < j {
				return false
			} else if start[i] == 'R' && j < i {
				return false
			}
		}
		i++
		j++
	}
	return i == j
}
