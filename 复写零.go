package main

import "fmt"

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	arr = []int{0, 1, 7, 6, 0, 2, 0, 7}
	arr = []int{8, 4, 5, 0, 0, 0, 0, 7}
	duplicateZeros(arr)
	fmt.Println(arr)
}

func duplicateZeros(arr []int) {
	n := len(arr)
	i, top := -1, 0
	for top < n {
		i++
		if arr[i] == 0 {
			top += 2
		} else {
			top++
		}
	}
	j := n - 1
	if top == n+1 {
		arr[j] = 0
		j--
		i--
	}
	for j >= 0 {
		arr[j] = arr[i]
		j--
		if arr[i] == 0 {
			arr[j] = 0
			j--
		}
		i--
	}
}
