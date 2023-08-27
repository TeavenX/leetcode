package main

import "math"

func main() {

}

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0)
	n := len(intervals)
	if n == 0 {
		ans = append(ans, newInterval)
		return ans
	}
	if newInterval[1] < intervals[0][0] {
		ans = append(ans, newInterval)
		ans = append(ans, intervals...)
		return ans
	} else if newInterval[0] > intervals[n-1][1] {
		ans = append(intervals, newInterval)
		return ans
	}
	ns, ne := newInterval[0], newInterval[1]
	for _, itv := range intervals {
		if itv[0] >= ns && itv[0] <= ne || itv[1] >= ns && itv[1] <= ne || itv[0] <= ns && itv[1] >= ne {
			ns = min(ns, itv[0])
			ne = max(ne, itv[1])
			continue
		}
		if ne < itv[0] {
			ans = append(ans, []int{ns, ne})
			ns, ne = math.MaxInt, math.MaxInt
		}
		ans = append(ans, itv)
	}
	if ne != math.MaxInt {
		ans = append(ans, []int{ns, ne})
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
