package main

import "sort"

func main() {

}

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	index := make([]int, n)
	for i := range index {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		ii, jj := index[i], index[j]
		if values[ii] == values[jj] {
			return labels[ii] < labels[jj]
		}
		return values[ii] > values[jj]
	})
	cache := make(map[int]int)
	ans := 0
	idx := 0
	for numWanted > 0 && idx < n {
		label := labels[index[idx]]
		if cache[label] < useLimit {
			ans += values[index[idx]]
			cache[label]++
			numWanted--
		}
		idx++
	}
	return ans
}
