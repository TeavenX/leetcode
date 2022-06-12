package main

import "fmt"

func main() {
	fmt.Println(consecutiveNumbersSum(77601076))
}

func consecutiveNumbersSumTimeOut(n int) int {
	result := 1
	left, right := 1, 1
	for right < n {
		sum := (left + right) * (right - left + 1)
		target := n << 1
		if sum == target {
			result++
			left++
			right++
		} else if sum < target {
			right++
		} else {
			left++
		}
	}
	return result
}

func consecutiveNumbersSum(n int) int {
	result, i := 0, 1
	for n > 0 {
		if n%i == 0 {
			result++
		}
		n -= i
		i++
	}
	return result
}
