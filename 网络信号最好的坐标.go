package main

import "math"

func main() {

}

func bestCoordinate(towers [][]int, radius int) []int {
	if len(towers) == 1 {
		if towers[0][2] == 0 {
			return []int{0, 0}
		}
		return towers[0][:2]
	}
	if len(towers) == 2 {
		if towers[0][2] > towers[1][2] {
			return towers[0][:2]
		}
		if towers[1][2] == 0 {
			return []int{0, 0}
		}
		return towers[1][:2]
	}
	xmin, ymin, xmax, ymax := towers[0][0], towers[0][1], towers[0][0], towers[0][1]
	for _, tower := range towers {
		if tower[0] < xmin {
			xmin = tower[0]
		} else if tower[0] > xmax {
			xmax = tower[0]
		}
		if tower[1] < ymin {
			ymin = tower[1]
		} else if tower[1] > ymax {
			ymax = tower[1]
		}
	}
	max := 0
	fr := float64(radius)
	mx, my := 0, 0
	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			sum := 0
			for _, tower := range towers {
				d := calcDistance(x, y, tower[0], tower[1])
				if d <= fr {
					r := int(float64(tower[2]) / (1 + d))
					sum += r
				}
			}
			if sum > max {
				max = sum
				mx = x
				my = y
			}
		}
	}
	return []int{mx, my}
}

func calcDistance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64(abs((x1-x2)*(x1-x2)) + abs((y1-y2)*(y1-y2))))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
