package main

import "math"

func main() {

}

func repairCars(ranks []int, cars int) int64 {
	mn := ranks[0]
	for _, x := range ranks {
		if x < mn {
			mn = x
		}
	}
	right := mn * cars * cars
	left := 0
	for left < right {
		mid := (left + right) >> 1
		sum := 0
		for _, r := range ranks {
			sum += int(math.Sqrt(float64(mid) / float64(r)))
		}
		if sum < cars {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return int64(right)
}
