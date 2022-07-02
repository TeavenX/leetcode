package main

import (
	"fmt"
	"strconv"
)

func main() {
	diffWaysToCompute("2*3-4*5")
}

func diffWaysToCompute(expression string) []int {
	if num, err := strconv.Atoi(expression); err == nil {
		return []int{num}
	}
	fmt.Println(expression)
	result := make([]int, 0)
	for idx, s := range expression {
		if s == '+' || s == '-' || s == '*' {
			left := diffWaysToCompute(expression[:idx])
			right := diffWaysToCompute(expression[idx+1:])
			for _, num1 := range left {
				for _, num2 := range right {
					total := 0
					switch s {
					case '+':
						total = num1 + num2
					case '-':
						total = num1 - num2
					case '*':
						total = num1 * num2
					}
					result = append(result, total)
				}
			}
		}
	}
	return result
}

var cache = map[string][]int{}

func diffWaysToComputeMemoization(expression string) []int {
	if num, err := strconv.Atoi(expression); err == nil {
		return []int{num}
	} else {
		if mem, ok := cache[expression]; ok {
			return mem
		}
	}
	fmt.Println(expression)
	result := make([]int, 0)
	for idx, s := range expression {
		if s == '+' || s == '-' || s == '*' {
			left := diffWaysToComputeMemoization(expression[:idx])
			right := diffWaysToComputeMemoization(expression[idx+1:])
			for _, num1 := range left {
				for _, num2 := range right {
					total := 0
					switch s {
					case '+':
						total = num1 + num2
					case '-':
						total = num1 - num2
					case '*':
						total = num1 * num2
					}
					result = append(result, total)
				}
			}
		}
	}
	cache[expression] = result
	return result
}
