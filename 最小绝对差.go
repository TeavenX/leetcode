package main

import "sort"

func main() {

}

func minimumAbsDifference(arr []int) [][]int {
	n := len(arr)
	if n == 2 {
		return [][]int{arr}
	}
	result := make([][]int, 0)
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < n; i++ {
		d := arr[i] - arr[i-1]
		if d < diff {
			diff = d
		}
	}
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] == diff {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}
