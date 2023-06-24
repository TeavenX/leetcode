package main

import "math"

func main() {

}

func calcDist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	corner := [4][2]int{{x1, y1}, {x1, y2}, {x2, y1}, {x2, y2}}
	for _, c := range corner {
		// 其中一个角在圆内
		if calcDist(xCenter, yCenter, c[0], c[1]) <= float64(radius) {
			return true
		}
	}
	// 圆心在矩形中
	if xCenter >= x1 && xCenter <= x2 && yCenter >= y1 && yCenter <= y2 {
		return true
	}
	// 平行于x轴的边与圆相交
	if xCenter >= x1 && xCenter <= x2 {
		if calcDist(xCenter, yCenter, xCenter, y1) <= float64(radius) ||
			calcDist(xCenter, yCenter, xCenter, y2) <= float64(radius) {
			return true
		}
	}
	// 垂直于x轴的边与圆相交
	if yCenter >= y1 && yCenter <= y2 {
		if calcDist(xCenter, yCenter, x1, yCenter) <= float64(radius) ||
			calcDist(xCenter, yCenter, x2, yCenter) <= float64(radius) {
			return true
		}
	}
	return false
}
