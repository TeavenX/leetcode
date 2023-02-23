package main

import "sort"

func main() {

}

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	minus := 0
	ans := 0
	for _, num := range nums {
		if d := num - minus; d > 0 {
			minus += d
			ans++
		}
	}
	return ans
}
