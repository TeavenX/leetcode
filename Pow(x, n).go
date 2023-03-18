package main

import "fmt"

func main() {
	fmt.Println(myPow(2, -2))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var ans float64 = 1
	for n > 0 {
		if n&1 == 1 {
			ans = ans * x
		}
		x *= x
		n >>= 1
	}
	return ans
}
