package main

import "sort"

func main() {

}

func minimumSize(nums []int, maxOperations int) int {
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return sort.Search(max, func(y int) bool {
		if y == 0 {
			return false
		}
		ops := 0
		for _, num := range nums {
			ops += (num - 1) / y
		}
		return ops <= maxOperations
	})
}
