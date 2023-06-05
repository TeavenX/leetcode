package main

import (
	"bytes"
	"strconv"
)

func main() {

}

func equalPairs(grid [][]int) int {
	n := len(grid)
	cache := make(map[string]int, n)
	ans := 0
	for _, row := range grid {
		t := bytes.Buffer{}
		for i := 0; i < n; i++ {
			t.WriteString(strconv.Itoa(row[i]))
			t.WriteByte('-')
		}
		t.WriteByte(';')
		cache[t.String()]++
	}
	for i := 0; i < n; i++ {
		t := bytes.Buffer{}
		for j := 0; j < n; j++ {
			t.WriteString(strconv.Itoa(grid[j][i]))
			t.WriteByte('-')
		}
		t.WriteByte(';')
		ans += cache[t.String()]
	}
	return ans
}
