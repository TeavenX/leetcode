package main

import "fmt"

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	//pushed := []int{2, 1, 0}
	//popped := []int{1, 2, 0}
	fmt.Println(validateStackSequencesV2(pushed, popped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	pIdx := 0
	for _, num := range pushed {
		if num != popped[pIdx] {
			stack = append(stack, num)
		} else {
			pIdx++
		}
	}
	fmt.Println(stack, pIdx)
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == popped[pIdx] {
			pIdx++
		} else {
			return false
		}
	}
	return true
}

func validateStackSequencesV2(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	pIdx := 0
	for _, num := range popped {
		if len(stack) > 0 && num == stack[len(stack)-1] {
			fmt.Println("pop", stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		} else {
			for pIdx < len(pushed) && num != pushed[pIdx] {
				stack = append(stack, pushed[pIdx])
				fmt.Println("push", pushed[pIdx])
				pIdx++
			}
			pIdx++
		}
	}
	return len(stack) == 0
}
