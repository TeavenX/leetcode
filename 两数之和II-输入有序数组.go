package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	search := func(left, right, target int) int {
		for left < right {
			mid := (left + right) >> 1
			if numbers[mid] == target {
				return mid
			} else if numbers[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return left
	}
	n := len(numbers)
	for i, num := range numbers {
		j := search(i+1, n-1, target-num)
		if numbers[j] == target-num {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}
