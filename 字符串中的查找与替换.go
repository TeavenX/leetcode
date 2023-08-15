package main

import "sort"

func main() {

}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	n := len(indices)
	index := make([]int, n)
	for i := range index {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		return indices[index[i]] > indices[index[j]]
	})
	for _, i := range index {
		idx := indices[i]
		ss := sources[i]
		if idx+len(ss) <= len(s) && s[idx:idx+len(ss)] == ss {
			s = s[:idx] + targets[i] + s[idx+len(ss):]
		}
	}
	return s
}
