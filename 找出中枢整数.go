package main

import "math"

func main() {

}

func pivotInteger(n int) int {
	for i := 1; i <= n; i++ {
		if (1+i)*i == (i+n)*(n-i+1) {
			return i
		}
	}
	return -1
}

func pivotInteger(n int) int {
	t := (n*n + n) / 2
	x := int(math.Sqrt(float64(t)))
	if x*x == t {
		return x
	}
	return -1
}
