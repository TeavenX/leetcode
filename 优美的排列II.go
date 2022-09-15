package main

import "fmt"

func main() {
	fmt.Println(constructArray(3, 1))
	fmt.Println(constructArray(3, 2))
	fmt.Println(constructArray(5, 4))
}

func constructArrayError(n int, k int) []int {
	result := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		result = append(result, i)
	}
	cur := result[len(result)-1]
	i := 0
	for len(result) < n {
		if k > 1 {
			result = append(result, n-i)
			i++
			k = (k + 1) / 2
		}
		cur++
		result = append(result, cur)
	}
	return result
}

func constructArrayErrorV2(n int, k int) []int {
	result := make([]int, 0, n)
	num := 1
	for len(result) < n {
		result = append(result, num)
		if num <= k/2 {
			result = append(result, n+1-num)
		}
		num++
	}
	return result
}

func constructArray(n int, k int) []int {
	result := make([]int, 0, n)
	result = append(result, 1)
	l, r := 2, n
	for i := 1; i < k; i++ {
		if i&1 == 0 {
			result = append(result, l)
			l++
		} else {
			result = append(result, r)
			r--
		}
	}
	for len(result) < n {
		if k&1 == 1 {
			result = append(result, l)
			l++
		} else {
			result = append(result, r)
			r--
		}
	}
	return result
}
