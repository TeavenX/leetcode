package main

import (
	"fmt"
	"strconv"
)

func main() {
	ret := addBinary("11", "10")
	ret = addBinary("1010", "1011")
	fmt.Println(ret)
	//fmt.Printf("%#v", strconv.Itoa(10))
}

func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	maxLen := max(lenA, lenB)
	result := ""
	carry := 0
	for i := 0; i < maxLen; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		result = strconv.Itoa(carry%2) + result
		carry /= 2
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
