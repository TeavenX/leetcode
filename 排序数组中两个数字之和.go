package main

import "fmt"

func main() {

	numbers, target := []int{1, 2, 4, 6, 10}, 8
	numbers, target = []int{-1, 0}, -1
	numbers, target = []int{-3, 3, 4, 90}, 0
	fmt.Println(twoSumV2(numbers, target))

}
func twoSum(numbers []int, target int) []int {

	lenNum := len(numbers)
	for i := 0; i < lenNum; i++ {
		for j := i + 1; j < lenNum; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumV2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{left, right}
}
