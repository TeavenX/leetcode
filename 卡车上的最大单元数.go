package main

import "sort"

func main() {

}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	count := 0
	for _, box := range boxTypes {
		if box[0] > truckSize {
			count += truckSize * box[1]
			break
		}
		count += box[0] * box[1]
		truckSize -= box[0]
	}
	return count
}

func maximumUnitsV2(boxTypes [][]int, truckSize int) int {
	cache := [1001]int{}
	for _, box := range boxTypes {
		cache[box[1]] += box[0]
	}
	count := 0
	for i := 1000; i > 0; i-- {
		if cache[i] == 0 {
			continue
		}
		if cache[i] >= truckSize {
			count += truckSize * i
			break
		}
		count += cache[i] * i
		truckSize -= cache[i]
	}
	return count
}
