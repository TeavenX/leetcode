package main

import "fmt"

func main() {
	//code := []int{5, 7, 1, 4}
	//k := 3
	code := []int{2, 4, 9, 3}
	k := -2
	fmt.Println(decryptV2(code, k))
}

func decrypt(code []int, k int) []int {
	if k == 0 {
		for i := range code {
			code[i] = 0
		}
		return code
	}
	n := len(code)
	sign := -1
	if k < 0 {
		sign = 1
	}
	result := make([]int, n)
	for i := range code {
		t := k
		s := 0
		for t != 0 {
			idx := (i + n + t) % n
			s += code[idx]
			t += sign
		}
		result[i] = s
	}
	return result
}

func decryptV2(code []int, k int) []int {
	if k == 0 {
		for i := range code {
			code[i] = 0
		}
		return code
	}
	n := len(code)
	sign := -1
	if k < 0 {
		sign = 1
	}
	result := make([]int, n)
	sum := 0
	t := k
	for t != 0 {
		sum += code[(n+t)%n]
		t += sign
	}
	for i := range code {
		head, tail := 0, 0
		if k > 0 {
			head = (i + k + 1) % n
			tail = (i + 1) % n
		} else {
			head = i
			tail = (n + i + k) % n
		}
		result[i] = sum
		sum += code[head]
		sum -= code[tail]
	}
	return result
}
