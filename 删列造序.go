package main

import "fmt"

func main() {
	strs := []string{"cba", "daf", "ghi"}
	fmt.Println(minDeletionSize(strs))
}

func minDeletionSize(strs []string) int {
	n, m := len(strs), len(strs[0])
	if n == 1 {
		return 0
	}
	result := 0
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if strs[j][i]-strs[j-1][i] > 122 {
				result++
				break
			}
		}
	}
	return result
}
