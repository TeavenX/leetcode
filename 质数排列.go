package main

import "fmt"

func main() {
	fmt.Println(numPrimeArrangements(5))
}

func numPrimeArrangements(n int) int {
	if n <= 2 {
		return 1
	}
	countNormal := 1
	countPrime := 1
	result := 1
	for i := 3; i <= n; i++ {
		if (i == 2 || i == 3 || i == 5 || i == 7) || (i%2 > 0 && i%3 > 0 && i%5 > 0 && i%7 > 0) {
			countPrime++
			result *= countPrime
		} else {
			countNormal++
			result *= countNormal
		}
		result %= 1e9 + 7
	}
	return result
}
