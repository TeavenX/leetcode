package main

import "sort"

func main() {

}

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	arr = arr[n/20 : n-n/20]
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return float64(sum) / float64(len(arr))
}
