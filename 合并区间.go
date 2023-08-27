package main

import "sort"

func main() {

}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			end = max(end, intervals[i][1])
			continue
		}
		ans = append(ans, []int{start, end})
		start, end = intervals[i][0], intervals[i][1]
	}
	ans = append(ans, []int{start, end})
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
