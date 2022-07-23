package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("  -+123"))
	fmt.Println(myAtoi("9223372036854775808"))
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := int64(0)
	flag := false
	sign := 1
	for _, str := range s {
		if str == ' ' {
			if flag {
				break
			}
			continue
		}
		num := str - '0'
		if !flag && (str == '+' || str == '-') {
			flag = true
			if str == '-' {
				sign = -1
			}
		} else if num <= 9 && num >= 0 {
			flag = true
			result = result*10 + int64(num)
			if result > math.MaxInt32 {
				break
			}
		} else {
			break
		}
	}
	if result > math.MaxInt32 || result < 0 {
		if sign == 1 {
			result = math.MaxInt32
		} else {
			result = -(math.MaxInt32 + 1)
		}
	} else {
		result = int64(sign) * result
	}
	return int(result)
}
