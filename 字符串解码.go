package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(decodeString("3[a2[c]]"))
	//fmt.Println(decodeString("3[a]2[bc]"))
	//fmt.Println(decodeString("10[leetcode]"))
	result := "bcbcxyxybcbcxyxyxdrtrtrtrt"
	r := decodeString("2[20[bc]31[xy]]xd4[rt]")
	fmt.Println(r)
	fmt.Println(result == r)
}

func decodeString(s string) string {
	var decode func(repeat int, str string) string
	decode = func(repeat int, str string) string {
		fmt.Println("in decode", repeat, str)
		base := []byte{}
		newR := -1
		startIdx := -1
		leftCount := 0
		for i := 0; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				if startIdx == -1 {
					temp, _ := strconv.Atoi(string(str[i]))
					if newR == -1 {
						newR = temp
					} else {
						newR = newR*10 + temp
					}
				}
			} else if str[i] == '[' {
				if startIdx == -1 {
					startIdx = i + 1
				}
				leftCount++
			} else if str[i] == ']' {
				leftCount--
				if leftCount == 0 {
					base = append(base, []byte(decode(newR, str[startIdx:i]))...)
					startIdx = -1
					newR = -1
				}
			} else if startIdx == -1 {
				base = append(base, str[i])
			}
		}
		ns := string(base)
		fmt.Println("before repeat", ns)
		for i := 1; i < repeat; i++ {
			ns += string(base)
		}
		fmt.Println("after repeat", ns)
		fmt.Println("============")
		return ns
	}
	return decode(1, s)
}
