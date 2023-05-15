package main

import "sort"

func main() {

}

func rearrangeBarcodes(barcodes []int) []int {
	cache := make(map[int]int)
	for _, num := range barcodes {
		cache[num]++
	}
	sort.Slice(barcodes, func(i, j int) bool {
		x, y := barcodes[i], barcodes[j]
		if cache[x] == cache[y] {
			return x < y
		}
		return cache[x] > cache[y]
	})
	n := len(barcodes)
	ans := make([]int, n)
	for j, k := 0, 0; k < 2; k++ {
		for i := k; i < n; i, j = i+2, j+1 {
			ans[i] = barcodes[j]
		}
	}
	return ans
}
