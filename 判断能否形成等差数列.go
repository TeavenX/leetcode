package main

import "sort"

func main() {

}

func canMakeArithmeticProgression(arr []int) bool {
	n := len(arr)
	if n == 2 {
		return true
	}
	sort.Ints(arr)
	dist := arr[1] - arr[0]
	for i := 1; i < n-1; i++ {
		if arr[i+1]-arr[i] != dist {
			return false
		}
	}
	return true
}
