package main

import "math"

func main() {

}

func nearestValidPoint(x int, y int, points [][]int) int {
	result, dist := -1, math.MaxInt
	for i := len(points) - 1; i >= 0; i-- {
		point := points[i]
		if point[0] == x || point[1] == y {
			d := abs(point[0] - x + point[1] - y)
			if d <= dist {
				result = i
				dist = d
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
