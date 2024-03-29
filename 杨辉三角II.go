package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}
func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			result[j] = result[j-1] + result[j]
		}
	}
	return result
}
