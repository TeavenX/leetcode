package main

import "sort"

func main() {

}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][1] < points[j][1] {
			return true
		} else if points[i][1] == points[j][1] && points[i][0] > points[j][0] {
			return true
		}
		return false
	})
	result := 1
	arrow := points[0][1]
	for _, point := range points {
		if point[0] <= arrow {
			continue
		} else {
			result++
			arrow = point[1]
		}
	}
	return result
}
