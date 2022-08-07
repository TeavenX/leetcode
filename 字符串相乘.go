package main

import "fmt"

func main() {
	fmt.Println(multiply("123", "456"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if num1 == "1" {
		return num2
	}
	if num2 == "1" {
		return num1
	}
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	totalLen := len(num1) + len(num2)
	result := make([]int, totalLen)
	base := 0
	for i := len(num1) - 1; i >= 0; i-- {
		carry := 0
		n1 := num1[i] - '0'
		cur := make([]int, len(num2)+1)
		idx := 0
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			temp := int(n1*n2) + carry
			carry = temp / 10
			cur[idx] = temp % 10
			idx++
		}
		for carry > 0 {
			cur[idx] = carry % 10
			carry = carry / 10
			idx++
		}
		idx, curIdx := 0, 0
		for k := 0; k < base; k++ {
			idx++
		}
		for curIdx < len(cur) {
			temp := cur[curIdx] + result[idx] + carry
			result[idx] = temp % 10
			carry = temp / 10
			idx++
			curIdx++
		}
		for carry > 0 {
			temp := result[idx] + carry
			result[idx] = temp % 10
			carry = temp / 10
			idx++
		}
		base++
	}
	ans := ""
	flag := false
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != 0 {
			flag = true
		}
		if flag {
			ans += string(byte(result[i] + '0'))
		}
	}
	return ans
}
