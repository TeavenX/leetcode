package main

import "math"

func main() {

}

func mctFromLeafValues(arr []int) int {
	n := len(arr)
	stack := make([]int, 0, n)
	stack = append(stack, math.MaxInt32)
	ans := 0
	for _, num := range arr {
		for stack[len(stack)-1] < num {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans += v * min(num, stack[len(stack)-1])
		}
		stack = append(stack, num)
	}
	for len(stack) > 2 {
		ans += stack[len(stack)-1] * stack[len(stack)-2]
		stack = stack[:len(stack)-1]
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
