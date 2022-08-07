package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(makeLargestSpecial("11011000"))
}

func makeLargestSpecial(s string) string {
	temp := make([]string, 0)
	count, last := 0, 0
	for i, b := range s {
		if b == '1' {
			count++
		} else {
			count--
		}
		if count == 0 {
			temp = append(temp, "1"+makeLargestSpecial(s[last+1:i])+"0")
			last = i + 1
		}
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] > temp[j]
	})
	return strings.Join(temp, "")
}
