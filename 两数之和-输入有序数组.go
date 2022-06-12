package main

import "fmt"

func main() {
	n, target := []int{2, 7, 11, 15}, 9
	//n, target = []int{2, 3, 4}, 6
	//n, target = []int{-1, 0}, -1
	//n, target = []int{-1, -1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, -2
	result := twoSumBest(n, target)
	fmt.Println(result)
}

func twoSum(numbers []int, target int) []int {
	left, n := 0, len(numbers)
	for left < n {
		right := left + 1
		for right < n {
			if numbers[left]+numbers[right] == target {
				return []int{left + 1, right + 1}
			}
			right++
		}
		left++
	}
	return []int{}
}

func twoSumBest(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}

func twoSumN167(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

func twoSum20220504(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
