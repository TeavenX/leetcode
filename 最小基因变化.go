package main

import "fmt"

func main() {
	start := "AACCGGTT"
	end := "AACCGGTA"
	bank := []string{"AACCGGTA"}

	start = "AACCTTGG"
	end = "AATTCCGG"
	bank = []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"}
	fmt.Println(minMutation(start, end, bank))
}

func hasOneDiff(a, b string) bool {
	diff := false
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if diff {
				return false
			}
			diff = true
		}
	}
	return diff
}

var dict = []string{"A", "T", "C", "G"}

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	cache := make(map[string]bool)
	for _, str := range bank {
		cache[str] = true
	}
	if _, exist := cache[end]; !exist {
		return -1
	}
	level := 0
	queue := []string{start}
	for len(queue) > 0 {
		qs := len(queue)
		for j := 0; j < qs; j++ {
			start = queue[0]
			queue = queue[1:]
			for i := 0; i < len(start); i++ {
				for _, str := range dict {
					temp := start[:i] + str + start[i+1:]
					if temp == start {
						continue
					}
					fmt.Println(temp)
					if temp == end {
						return level + 1
					}
					if exist := cache[temp]; exist {
						cache[temp] = false
						queue = append(queue, temp)
					}
				}
			}
		}
		level++
	}
	return -1
}
