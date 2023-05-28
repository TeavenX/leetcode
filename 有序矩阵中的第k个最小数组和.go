package main

import "sort"

func main() {

}

func kthSmallest(mat [][]int, k int) int {
	a := []int{0}
	for _, row := range mat {
		b := make([]int, 0, len(a)*len(row))
		for _, v := range row {
			for _, t := range a {
				b = append(b, v+t)
			}
		}
		sort.Ints(b)
		if len(b) > k {
			b = b[:k]
		}
		a = b
	}
	return a[k-1]
}
