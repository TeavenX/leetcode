package main

import "sort"

func main() {

}

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] < intervals[j][1] {
			return true
		} else if intervals[i][1] == intervals[j][1] && intervals[i][0] > intervals[j][0] {
			return true
		}
		return false
	})
	result := 2
	preEnd, end := intervals[0][1]-1, intervals[0][1]
	for _, interval := range intervals {
		// x <= preEnd
		if interval[0] <= preEnd {
			continue
		}
		// x <= end
		if interval[0] <= end {
			result++
			preEnd = end
			end = interval[1]
		} else {
			// end < x
			result += 2
			preEnd = interval[1] - 1
			end = interval[1]
		}
	}
	return result
}
