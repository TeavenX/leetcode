package main

import "fmt"

func main() {
	fmt.Println(minFlipsMonoIncr("00110"))
	fmt.Println(minFlipsMonoIncr("00011000"))
	fmt.Println(minFlipsMonoIncr("100000001010000"))
}

func minFlipsMonoIncr(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	dp0, dp1 := 0, 0
	for _, str := range s {
		dp1 = min(dp0, dp1)
		if str == '1' {
			dp0++
		} else {
			dp1++
		}
	}
	return min(dp0, dp1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
