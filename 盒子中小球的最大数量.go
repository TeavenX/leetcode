package main

import "fmt"

func main() {
	//for i := int(1e3); i < 1e5; i++ {
	//    fmt.Printf("%d ", calc(i))
	//}
	fmt.Println(calc(int(1e5) - 1))
}

func countBalls(lowLimit int, highLimit int) int {
	//cache := make(map[int]int)
	cache := [46]int{}
	max := 0
	for i := lowLimit; i <= highLimit; i++ {
		k := calc(i)
		cache[k]++
		if cache[k] > max {
			max = cache[k]
		}
	}
	return max
}

func calc(a int) int {
	r := 0
	for a > 0 {
		r += a % 10
		a /= 10
	}
	return r
}
